package model

type ContainerRequest struct {
	ContainerName string `json:"containerName"`
}

type BlobContainerListResponse struct {
	BlobName []string `json:"blobName"`
}
