package service

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Siravitt/azure-storage/model"
)

func (s service) BlobContainerList(ctx context.Context, req model.ContainerRequest) (*model.BlobContainerListResponse, error) {
	pager := s.client.NewListBlobsFlatPager(req.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	respSrv := model.BlobContainerListResponse{
		BlobName: []string{},
	}

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Panicf("read blob file error: %s", err.Error())
			return nil, err
		}

		for _, blob := range resp.Segment.BlobItems {
			respSrv.BlobName = append(respSrv.BlobName, *blob.Name)
		}
	}

	return &respSrv, nil
}
