package types

import (
	"fmt"

	"github.com/multycloud/multy/api/errors"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/flags"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/route_table_association"
	"github.com/multycloud/multy/resources/output/subnet"
	"github.com/multycloud/multy/util"
	"github.com/multycloud/multy/validate"
)

/*
Notes:
Azure: New subnets will be associated with a default route table to block traffic to internet
*/

var subnetMetadata = resources.ResourceMetadata[*resourcespb.SubnetArgs, *Subnet, *resourcespb.SubnetResource]{
	ImportFunc:        NewSubnet,
	CreateFunc:        CreateSubnet,
	UpdateFunc:        UpdateSubnet,
	ReadFromStateFunc: SubnetFromState,
	ExportFunc: func(r *Subnet, _ *resources.Resources) (*resourcespb.SubnetArgs, bool, error) {
		return r.Args, true, nil
	},
	AbbreviatedName: "vn",
}

type Subnet struct {
	resources.ChildResourceWithId[*VirtualNetwork, *resourcespb.SubnetArgs]

	VirtualNetwork *VirtualNetwork
}

func CreateSubnet(resourceId string, args *resourcespb.SubnetArgs, others *resources.Resources) (*Subnet, error) {
	return NewSubnet(resourceId, args, others)
}

func UpdateSubnet(resource *Subnet, vn *resourcespb.SubnetArgs, others *resources.Resources) error {
	resource.Args = vn
	return nil
}

func NewSubnet(resourceId string, subnet *resourcespb.SubnetArgs, others *resources.Resources) (*Subnet, error) {
	s := &Subnet{
		ChildResourceWithId: resources.ChildResourceWithId[*VirtualNetwork, *resourcespb.SubnetArgs]{
			ResourceId: resourceId,
			Args:       subnet,
		},
	}
	vn, err := resources.Get[*VirtualNetwork](resourceId, others, subnet.VirtualNetworkId)
	if err != nil {
		return nil, errors.ValidationErrors([]validate.ValidationError{s.NewValidationError(err, "virtual_network_id")})
	}
	s.Parent = vn
	s.VirtualNetwork = vn
	return s, nil
}

func SubnetFromState(r *Subnet, state *output.TfState) (*resourcespb.SubnetResource, error) {
	if flags.DryRun {
		return &resourcespb.SubnetResource{
			CommonParameters: &commonpb.CommonChildResourceParameters{
				ResourceId:  r.ResourceId,
				NeedsUpdate: false,
			},
			Name:             r.Args.Name,
			CidrBlock:        r.Args.CidrBlock,
			AvailabilityZone: r.Args.AvailabilityZone,
			VirtualNetworkId: r.Args.VirtualNetworkId,
		}, nil
	}
	out := new(resourcespb.SubnetResource)
	out.CommonParameters = &commonpb.CommonChildResourceParameters{
		ResourceId:  r.ResourceId,
		NeedsUpdate: false,
	}
	out.AvailabilityZone = r.Args.AvailabilityZone
	out.VirtualNetworkId = r.Args.GetVirtualNetworkId()

	id, err := resources.GetMainOutputRef(r)
	if err != nil {
		return nil, err
	}
	switch r.GetCloud() {
	case common.AWS:
		stateResource, err := output.GetParsed[subnet.AwsSubnet](state, id)
		if err != nil {
			return nil, err
		}
		out.Name = stateResource.AwsResource.Tags["Name"]
		out.CidrBlock = stateResource.CidrBlock
	case common.AZURE:
		stateResource, err := output.GetParsed[subnet.AzureSubnet](state, id)
		if err != nil {
			return nil, err
		}
		out.Name = stateResource.Name
		out.CidrBlock = stateResource.AddressPrefixes[0]
	}

	return out, nil
}

func (r *Subnet) GetMetadata() resources.ResourceMetadataInterface {
	return &subnetMetadata
}

func (r *Subnet) Translate(ctx resources.MultyContext) ([]output.TfBlock, error) {
	if r.GetCloud() == commonpb.CloudProvider_AWS {
		location := r.VirtualNetwork.GetLocation()
		az, err := common.GetAvailabilityZone(location, int(r.Args.AvailabilityZone), r.GetCloud())
		if err != nil {
			return nil, err
		}
		awsSubnet := subnet.AwsSubnet{
			AwsResource:      common.NewAwsResource(r.ResourceId, r.Args.Name),
			CidrBlock:        r.Args.CidrBlock,
			VpcId:            r.VirtualNetwork.GetVirtualNetworkId(),
			AvailabilityZone: az,
		}
		// This flag needs to be set so that eks nodes can connect to the kubernetes cluster
		// https://aws.amazon.com/blogs/containers/upcoming-changes-to-ip-assignment-for-eks-managed-node-groups/
		// How to tell if this subnet is private?
		if len(resources.GetAllResourcesWithRef(ctx, func(k *KubernetesNodePool) *Subnet { return k.Subnet }, r)) > 0 {
			awsSubnet.MapPublicIpOnLaunch = true
		}
		if len(resources.GetAllResourcesWithRef(ctx, func(k *KubernetesCluster) *Subnet { return k.DefaultNodePool.Subnet }, r)) > 0 {
			awsSubnet.MapPublicIpOnLaunch = true
		}
		return []output.TfBlock{awsSubnet}, nil
	} else if r.GetCloud() == commonpb.CloudProvider_AZURE {
		var azResources []output.TfBlock
		azSubnet := subnet.AzureSubnet{
			AzResource: &common.AzResource{
				TerraformResource: output.TerraformResource{ResourceId: r.ResourceId},
				Name:              r.Args.Name,
				ResourceGroupName: GetResourceGroupName(r.VirtualNetwork.Args.GetCommonParameters().GetResourceGroupId()),
			},
			AddressPrefixes:    []string{r.Args.CidrBlock},
			VirtualNetworkName: r.VirtualNetwork.GetVirtualNetworkName(),
		}
		azSubnet.ServiceEndpoints = getServiceEndpointSubnetReferences(ctx, r)
		azResources = append(azResources, azSubnet)

		// there must be a better way to do this
		if !checkSubnetRouteTableAssociated(ctx, r) {
			rt, err := r.VirtualNetwork.GetAssociatedRouteTableId()
			if err != nil {
				return nil, err
			}
			subnetId, err := resources.GetMainOutputId(r)
			if err != nil {
				return nil, err
			}
			rtAssociation := route_table_association.AzureRouteTableAssociation{
				AzResource: &common.AzResource{
					TerraformResource: output.TerraformResource{ResourceId: r.ResourceId},
				},
				SubnetId:     subnetId,
				RouteTableId: rt,
			}
			azResources = append(azResources, rtAssociation)
		}

		return azResources, nil
	}
	return nil, fmt.Errorf("cloud %r is not supported for this resource type ", r.GetCloud().String())
}

func (r *Subnet) GetId() string {
	t, _ := r.GetMainResourceName()
	return fmt.Sprintf("%s.%s.id", t, r.ResourceId)
}

func getServiceEndpointSubnetReferences(ctx resources.MultyContext, r *Subnet) []string {
	const (
		DATABASE = "Microsoft.Sql"
	)

	serviceEndpoints := map[string]bool{}
	if len(resources.GetAllResourcesWithListRef(ctx, func(db *Database) []*Subnet { return db.Subnets }, r)) > 0 {
		serviceEndpoints[DATABASE] = true
	}
	return util.SortedKeys(serviceEndpoints)
}

func checkSubnetRouteTableAssociated(ctx resources.MultyContext, r *Subnet) bool {
	return len(resources.GetAllResourcesWithRef(ctx, func(rt *RouteTableAssociation) *Subnet { return rt.Subnet }, r)) > 0
}

func (r *Subnet) Validate(ctx resources.MultyContext) (errs []validate.ValidationError) {
	//if vn.Name contains not letters,numbers,_,- { return false }
	//if vn.Name length? { return false }
	//if vn.CidrBlock valid CIDR { return false }
	//if vn.AvailbilityZone valid { return false }
	if len(r.Args.CidrBlock) == 0 { // max len?
		errs = append(errs, r.NewValidationError(fmt.Errorf("%r cidr_block length is invalid", r.ResourceId), "cidr_block"))
	}

	return errs
}

func (r *Subnet) GetMainResourceName() (string, error) {
	switch r.GetCloud() {
	case common.AWS:
		return subnet.AwsResourceName, nil
	case common.AZURE:
		return subnet.AzureResourceName, nil
	default:
		return "", fmt.Errorf("unknown cloud %r", r.GetCloud().String())
	}
}

func (r *Subnet) GetCloudSpecificLocation() string {
	return r.VirtualNetwork.GetCloudSpecificLocation()
}
