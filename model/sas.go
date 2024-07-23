package model

type GenerateSASRequest struct {
	ContentType   string `json:"contentType"`
	ContainerName string `json:"containerName"`
	FileName      string `json:"fileName"`
}

type GenerateSASResponse struct {
	SASURL string `json:"sasURL"`
}
