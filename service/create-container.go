package service

import (
	"context"

	"github.com/Siravitt/azure-storage/model"
	"github.com/labstack/gommon/log"
)

func (s service) CreateContainer(ctx context.Context, req model.ContainerRequest) error {
	_, err := s.client.CreateContainer(ctx, req.ContainerName, nil)
	if err != nil {
		log.Errorf("CreateContainer error: %s", err)
		return err
	}

	return nil
}
