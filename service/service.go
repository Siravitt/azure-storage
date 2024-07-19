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
	GenerateSAS(req model.GenerateSASRequest) (*model.GenerateSASResponse, error)
	CreateContainer(ctx context.Context, containerName string) error
	DeleteContainer(ctx context.Context, containerName string) error
}
