package service

import (
	"bytes"
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Siravitt/azure-storage/model"
)

func (s service) BlobDownload(ctx context.Context, req model.BlobDownloadRequest) (*bytes.Buffer, error) {
	get, err := s.client.DownloadStream(ctx, req.ContainerName, req.FileName, nil)
	if err != nil {
		log.Panicf("DownloadStream error: %s", err)
		return nil, err
	}

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(context.TODO(), &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	if err != nil {
		log.Panicf("DownloadStream error: %s", err)
		return nil, err
	}

	err = retryReader.Close()
	if err != nil {
		log.Panicf("DownloadStream error: %s", err)
		return nil, err
	}

	return &downloadedData, nil
}
