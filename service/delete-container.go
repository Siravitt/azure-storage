package service

import (
	"context"
	"log"

	"github.com/Siravitt/azure-storage/model"
)

func (s service) DeleteContainer(ctx context.Context, req model.DeleteContainerRequest) error {
	_, err := s.client.DeleteContainer(ctx, req.ContainerName, nil)
	if err != nil {
		log.Fatalf("DeleteContainer error: %s", err)
		return err
	}

	return nil
}
