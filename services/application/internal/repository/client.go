package repository

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/lucas-hill/credit-decision/services/application/internal/config"
)

func NewDynamoDBClient(ctx context.Context, cfg *config.Config) (*dynamodb.Client, error) {
	var opts []func(*awsconfig.LoadOptions) error

	opts = append(opts, awsconfig.WithRegion(cfg.AWSRegion))

	if cfg.IsLocal() {
		// Local development: dummy credentials, custom endpoint
		opts = append(opts, awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("local", "local", ""),
		))
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(awsCfg, func(o *dynamodb.Options) {
		if cfg.DynamoDBEndpoint != "" {
			o.BaseEndpoint = &cfg.DynamoDBEndpoint
		}
	})

	return client, nil
}
