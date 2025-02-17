package db

import (
	"fmt"
	aws_client "github.com/multycloud/multy/api/aws"
	"github.com/multycloud/multy/api/errors"
	"github.com/multycloud/multy/api/proto/configpb"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

type userConfigStorage struct {
	AwsClient aws_client.AwsClient
}

func (d *userConfigStorage) StoreUserConfig(config *configpb.Config, lock *ConfigLock) error {
	if !lock.IsActive() {
		return fmt.Errorf("unable to store user config because lock is invalid")
	}
	log.Printf("[INFO] Storing user config from api_key %s\n", config.UserId)
	b, err := protojson.Marshal(config)
	if err != nil {
		return err
	}

	err = d.AwsClient.SaveFile(config.UserId, configFile, string(b))
	if err != nil {
		return errors.InternalServerErrorWithMessage("error storing configuration", err)
	}
	return nil
}

func (d *userConfigStorage) LoadUserConfig(userId string, lock *ConfigLock) (*configpb.Config, error) {
	if lock != nil && !lock.IsActive() {
		return nil, fmt.Errorf("unable to load user config because lock is invalid")
	}
	log.Printf("[INFO] Loading config from api_key %s\n", userId)
	result := configpb.Config{
		UserId: userId,
	}

	tfFileStr, err := d.AwsClient.ReadFile(userId, configFile)
	if err != nil {
		return nil, errors.InternalServerErrorWithMessage("error reading configuration", err)
	}

	if tfFileStr != "" {

		err := protojson.Unmarshal([]byte(tfFileStr), &result)
		if err != nil {
			return nil, errors.InternalServerErrorWithMessage("error parsing configuration", err)
		}
	}
	return &result, nil
}

func newUserConfigStorage(awsClient aws_client.AwsClient) (*userConfigStorage, error) {

	return &userConfigStorage{
		AwsClient: awsClient,
	}, nil
}
