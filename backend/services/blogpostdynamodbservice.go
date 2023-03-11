package services

import (
	"context"
	"errors"

	"blog.jonastrogen.se/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BlogpostDynamoDBService struct {
	db    *dynamodb.Client
	table string
}

func NewBlogpostDynamoDBService(table string, cfg aws.Config) *BlogpostDynamoDBService {
	return &BlogpostDynamoDBService{
		db:    dynamodb.NewFromConfig(cfg),
		table: table,
	}
}

func (s *BlogpostDynamoDBService) GetBlogpost(ctx context.Context, name string) (*models.BlogPost, error) {
	titleAttr, err := attributevalue.Marshal(name)
	if err != nil {
		return nil, err
	}
	result, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]types.AttributeValue{
			"title": titleAttr,
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		msg := "Could not find '" + name + "'"
		return nil, errors.New(msg)
	}

	blogpost := models.BlogPost{}

	err = attributevalue.UnmarshalMap(result.Item, &blogpost)
	if err != nil {
		return nil, err
	}

	return &blogpost, nil
}

func (s *BlogpostDynamoDBService) GetMetadata(ctx context.Context, key string) (*models.Metadata, error) {
	titleAttr, err := attributevalue.Marshal(key)
	if err != nil {
		return nil, err
	}
	result, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]types.AttributeValue{
			"title": titleAttr,
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		msg := "Could not find '" + key + "'"
		return nil, errors.New(msg)
	}

	metadata := models.Metadata{}

	err = attributevalue.UnmarshalMap(result.Item, &metadata)
	if err != nil {
		return nil, err
	}

	// FIXME
	metadata.Key = metadata.Title

	return &metadata, nil
}

// WIP fix this
func (s *BlogpostDynamoDBService) ListMetadata(ctx context.Context) (*[]models.Metadata, error) {
	// FIXME
	res, err := s.db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(s.table),
		// ExclusiveStartKey:         map[string]types.AttributeValue{},
		// ExpressionAttributeNames:  map[string]string{},
		// ExpressionAttributeValues: map[string]types.AttributeValue{},
		// FilterExpression:          new(string),
		Limit: aws.Int32(10),
	})
	if err != nil {
		return nil, err
	}

	if res.Items == nil {
		return nil, errors.New("could not find any items")
	}

	list := make([]models.Metadata, len(res.Items))
	for i, m := range res.Items {
		metadata := models.Metadata{}

		err = attributevalue.UnmarshalMap(m, &metadata)
		if err != nil {
			return nil, err
		}
		// FIXME
		metadata.Key = metadata.Title

		list[i] = metadata
	}

	return &list, nil
}
