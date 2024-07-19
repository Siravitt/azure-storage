package service

import (
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
	CreateContainer(ctx context.Context, req model.CreateContainerRequest) error
	DeleteContainer(ctx context.Context, req model.DeleteContainerRequest) error
}
