package credential

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Credential struct {
	awsConfig aws.Config
}

func NewCredential(awsConfig AWSConfig, ctx context.Context) (*Credential, error) {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsConfig.Region))
	if err != nil {
		return nil, err
	}

	// Process environment variables into the AWS config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err)
	}

	// Create a new credential cache
	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		awsConfig.AccessKeyID, awsConfig.SecretAccessKey, ""))

	// Assign the credential to the AWS config
	cfg.Credentials = creds

	return &Credential{awsConfig: cfg}, nil
}

func (c *Credential) GetAWSConfig() aws.Config {
	return c.awsConfig
}
