// Package storage provides Cloud Storage client integration
package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/mikouaj/go-rest-cloud-storage/internal/api"
	"google.golang.org/api/iterator"
)

type StorageObject struct {
	Name string
}

type StorageClient interface {
	ListObjects(bucket string) (api.StorageObjects, error)
}

type cloudStorageClient struct {
	ctx    context.Context
	client *storage.Client
}

func NewStorageClient(ctx context.Context) (StorageClient, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &cloudStorageClient{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c cloudStorageClient) ListObjects(bucketName string) (api.StorageObjects, error) {
	results := make(api.StorageObjects, 0)
	query := &storage.Query{}
	bucket := c.client.Bucket(bucketName)
	it := bucket.Objects(c.ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := api.StorageObject{
			Name: attrs.Name,
		}
		results = append(results, result)
	}
	return results, nil
}
