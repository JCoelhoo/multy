package types

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/flags"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/public_ip"
	"github.com/multycloud/multy/validate"
)

/*
Notes
AWS: NIC association done on public_ip
Azure: NIC association done on NIC creation
*/

var publicIpMetadata = resources.ResourceMetadata[*resourcespb.PublicIpArgs, *PublicIp, *resourcespb.PublicIpResource]{
	CreateFunc:        CreatePublicIp,
	UpdateFunc:        UpdatePublicIp,
	ReadFromStateFunc: PublicIpFromState,
	ExportFunc: func(r *PublicIp, _ *resources.Resources) (*resourcespb.PublicIpArgs, bool, error) {
		return r.Args, true, nil
	},
	ImportFunc:      NewPublicIp,
	AbbreviatedName: "pip",
}

type PublicIp struct {
	resources.ResourceWithId[*resourcespb.PublicIpArgs]
}

func (r *PublicIp) GetMetadata() resources.ResourceMetadataInterface {
	return &publicIpMetadata
}

func CreatePublicIp(resourceId string, args *resourcespb.PublicIpArgs, others *resources.Resources) (*PublicIp, error) {
	if args.CommonParameters.ResourceGroupId == "" {
		// todo: maybe put in the same RG as VM?
		rgId, err := NewRg("pip", others, args.GetCommonParameters().GetLocation(), args.GetCommonParameters().GetCloudProvider())
		if err != nil {
			return nil, err
		}
		args.CommonParameters.ResourceGroupId = rgId
	}

	return NewPublicIp(resourceId, args, others)
}

func UpdatePublicIp(resource *PublicIp, vn *resourcespb.PublicIpArgs, others *resources.Resources) error {
	resource.Args = vn
	return nil
}

func PublicIpFromState(resource *PublicIp, state *output.TfState) (*resourcespb.PublicIpResource, error) {
	var err error
	ip := "dryrun"
	if !flags.DryRun {
		ip, err = getIp(resource.ResourceId, state, resource.Args.CommonParameters.CloudProvider)
		if err != nil {
			return nil, err
		}
	}

	return &resourcespb.PublicIpResource{
		CommonParameters: &commonpb.CommonResourceParameters{
			ResourceId:      resource.ResourceId,
			ResourceGroupId: resource.Args.CommonParameters.ResourceGroupId,
			Location:        resource.Args.CommonParameters.Location,
			CloudProvider:   resource.Args.CommonParameters.CloudProvider,
			NeedsUpdate:     false,
		},
		Name: resource.Args.Name,

		Ip: ip,
	}, nil
}

func NewPublicIp(resourceId string, args *resourcespb.PublicIpArgs, others *resources.Resources) (*PublicIp, error) {
	return &PublicIp{
		ResourceWithId: resources.ResourceWithId[*resourcespb.PublicIpArgs]{
			ResourceId: resourceId,
			Args:       args,
		},
	}, nil
}

func (r *PublicIp) Translate(resources.MultyContext) ([]output.TfBlock, error) {
	if r.GetCloud() == commonpb.CloudProvider_AWS {
		return []output.TfBlock{
			public_ip.AwsElasticIp{
				AwsResource: common.NewAwsResource(r.ResourceId, r.Args.Name),
				//Vpc:        true,
			},
		}, nil
	} else if r.GetCloud() == commonpb.CloudProvider_AZURE {
		return []output.TfBlock{
			public_ip.AzurePublicIp{
				AzResource: common.NewAzResource(
					r.ResourceId, r.Args.Name, GetResourceGroupName(r.Args.GetCommonParameters().ResourceGroupId),
					r.GetCloudSpecificLocation(),
				),
				AllocationMethod: "Static",
			},
		}, nil
	}
	return nil, fmt.Errorf("cloud %s is not supported for this resource type ", r.GetCloud().String())
}

func (r *PublicIp) GetId(cloud commonpb.CloudProvider) string {
	types := map[commonpb.CloudProvider]string{common.AWS: public_ip.AwsResourceName, common.AZURE: public_ip.AzureResourceName}
	return fmt.Sprintf("%s.%s.id", types[cloud], r.ResourceId)
}

func getIp(resourceId string, state *output.TfState, cloud commonpb.CloudProvider) (string, error) {
	switch cloud {
	case commonpb.CloudProvider_AWS:
		values, err := state.GetValues(public_ip.AwsElasticIp{}, resourceId)
		if err != nil {
			return "", err
		}
		return values["public_ip"].(string), nil
	case commonpb.CloudProvider_AZURE:
		values, err := state.GetValues(public_ip.AzurePublicIp{}, resourceId)
		if err != nil {
			return "", err
		}
		return values["ip_address"].(string), nil
	}

	return "", fmt.Errorf("unknown cloud: %s", cloud.String())
}

func (r *PublicIp) Validate(ctx resources.MultyContext) (errs []validate.ValidationError) {
	errs = append(errs, r.ResourceWithId.Validate()...)
	return errs
}

func (r *PublicIp) GetMainResourceName() (string, error) {
	switch r.GetCloud() {
	case commonpb.CloudProvider_AWS:
		return public_ip.AwsResourceName, nil
	case commonpb.CloudProvider_AZURE:
		return public_ip.AzureResourceName, nil
	default:
		return "", fmt.Errorf("unknown cloud %s", r.GetCloud().String())
	}
}
