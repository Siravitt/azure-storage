package model

type BlobDownloadRequest struct {
	ContainerName string `json:"containerName"`
	FileName      string `json:"fileName"`
}
