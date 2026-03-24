package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/lucas-hill/credit-decision/services/application/internal/model"
)

type DynamoDBRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBRepository(client *dynamodb.Client, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{
		client:    client,
		tableName: tableName,
	}
}

func (r *DynamoDBRepository) Create(ctx context.Context, app *model.Application) error {
	app.PK = model.BuildPK(app.ID)
	app.SK = model.ApplicationSKMetadata

	item, err := attributevalue.MarshalMap(app)
	if err != nil {
		return fmt.Errorf("failed to marshal application: %w", err)
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           &r.tableName,
		Item:                item,
		ConditionExpression: strPtr("attribute_not_exists(PK)"),
	})
	if err != nil {
		return fmt.Errorf("failed to put application: %w", err)
	}

	return nil
}

func (r *DynamoDBRepository) GetByID(ctx context.Context, id string) (*model.Application, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: model.BuildPK(id)},
			"SK": &types.AttributeValueMemberS{Value: model.ApplicationSKMetadata},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get application: %w", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	var app model.Application
	if err := attributevalue.UnmarshalMap(result.Item, &app); err != nil {
		return nil, fmt.Errorf("failed to unmarshal application: %w", err)
	}

	return &app, nil
}

func strPtr(s string) *string {
	return &s
}
