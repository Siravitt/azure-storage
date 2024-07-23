package service

import (
	"bytes"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Siravitt/azure-storage/model"
	"github.com/Siravitt/azure-storage/repository"
)

type service struct {
	repo   repository.Repository
	client *azblob.Client
}

func NewService(repo repository.Repository, client *azblob.Client) Service {
	return service{
		repo:   repo,
		client: client,
	}
}

type Service interface {
	GenerateSASUpload(req model.GenerateSASRequest) (*model.GenerateSASResponse, error)
	GenerateSASRead(req model.GenerateSASRequest) (*model.GenerateSASResponse, error)
	CreateContainer(ctx context.Context, req model.ContainerRequest) error
	DeleteContainer(ctx context.Context, req model.ContainerRequest) error
	BlobContainerList(ctx context.Context, req model.ContainerRequest) (*model.BlobContainerListResponse, error)
	BlobDownload(ctx context.Context, req model.BlobDownloadRequest) (*bytes.Buffer, error)
}
