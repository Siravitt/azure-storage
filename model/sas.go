package model

type GenerateSASRequest struct {
	ContainerName string `json:"containerName"`
	FileName      string `json:"fileName"`
}

type GenerateSASResponse struct {
	SASURL string `json:"sasURL"`
}
