package model

type CreateContainerRequest struct {
	ContainerName string `json:"containerName"`
}

type DeleteContainerRequest struct {
	ContainerName string `json:"containerName"`
}
