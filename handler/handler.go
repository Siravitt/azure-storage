package handler

import (
	"github.com/Siravitt/azure-storage/service"
	"github.com/labstack/echo/v4"
)

type handler struct {
	srv service.Service
}

func NewHandler(srv service.Service) Handler {
	return handler{srv: srv}
}

type Handler interface {
	Health(c echo.Context) error
	GenerateSASUpload(c echo.Context) error
	GenerateSASRead(c echo.Context) error
	CreateContainer(c echo.Context) error
	DeleteContainer(c echo.Context) error
}
