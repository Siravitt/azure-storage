package service

import "github.com/Siravitt/azure-storage/repository"

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}

type Service interface {
}
