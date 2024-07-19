package service

import (
	"context"

	"github.com/labstack/gommon/log"
)

func (s service) CreateContainer(ctx context.Context, containerName string) error {
	_, err := s.client.CreateContainer(ctx, containerName, nil)
	if err != nil {
		log.Errorf("CreateContainer error: %s", err)
		return err
	}

	return nil
}
