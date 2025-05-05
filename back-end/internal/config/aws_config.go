package config

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AwsConfig struct {
	CognitoClient *cognitoidentityprovider.Client
}

func NewAwsConfig() *AwsConfig {

	return &AwsConfig{}
}
