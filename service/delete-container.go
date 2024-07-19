package service

import (
	"context"
	"log"
)

func (s service) DeleteContainer(ctx context.Context, containerName string) error {
	_, err := s.client.DeleteContainer(ctx, containerName, nil)
	if err != nil {
		log.Fatalf("DeleteContainer error: %s", err)
		return err
	}

	return nil
}
