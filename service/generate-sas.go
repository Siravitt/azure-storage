package service

import (
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Siravitt/azure-storage/model"
)

func (s service) GenerateSAS(req model.GenerateSASRequest) (*model.GenerateSASResponse, error) {
	containerClient := s.client.ServiceClient().NewContainerClient(req.ContainerName)
	blobClient := containerClient.NewBlobClient(req.FileName)

	permission := sas.BlobPermissions{
		Read:   true,
		Add:    true,
		Create: true,
		Write:  true,
	}

	expiryTime := time.Now().Add(24 * time.Hour)
	sasURL, err := blobClient.GetSASURL(permission, expiryTime, nil)
	if err != nil {
		log.Panicf("GetSASURL error: %s", err)
		return nil, err
	}

	resp := model.GenerateSASResponse{
		SASURL: sasURL,
	}

	return &resp, nil
}
