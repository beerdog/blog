package services

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"blog.jonastrogen.se/models"
)

type BlogpostS3Service struct {
	bucket   string
	s3Client *s3.Client
}

func NewBlogpostS3Service(bucket string, cfg aws.Config) *BlogpostS3Service {
	return &BlogpostS3Service{
		bucket:   bucket,
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (s *BlogpostS3Service) GetBlogpost(ctx context.Context, name string) (*models.BlogPost, error) {
	output, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String("blogposts/" + name + ".md"),
	})

	if err != nil {
		return nil, err
	}

	s3objectBytes, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	metadata, err := s.GetMetadata(ctx, name)
	if err != nil {
		return nil, err
	}

	blogPost := models.BlogPost{
		Content:  string(s3objectBytes[:]), // TODO suboptimal, should fix serializing dynamodb content to byte[] instead of string.
		Metadata: *metadata,
	}

	return &blogPost, nil
}

func (s *BlogpostS3Service) getMetadataByKey(ctx context.Context, key string) (*models.Metadata, error) {
	output, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	s3objectBytes, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	metadata := models.Metadata{}

	err = json.Unmarshal(s3objectBytes, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (s *BlogpostS3Service) GetMetadata(ctx context.Context, key string) (*models.Metadata, error) {
	return s.getMetadataByKey(ctx, "blogposts/"+key+".json")
}

func (s *BlogpostS3Service) ListMetadata(ctx context.Context) (*[]models.Metadata, error) {
	output, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String("blogposts/"),
	})

	if err != nil {
		return nil, err
	}

	metadataList := []models.Metadata{}
	for _, file := range output.Contents {
		if strings.HasSuffix(*file.Key, ".json") {
			metadata, err := s.getMetadataByKey(ctx, *file.Key)
			if err != nil {
				return nil, err
			}
			metadataList = append(metadataList, *metadata)
		}
	}

	return &metadataList, err
}
