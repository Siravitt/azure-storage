package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Siravitt/azure-storage/model"
	"github.com/labstack/echo/v4"
)

func (h handler) CreateContainer(c echo.Context) error {
	ctx := context.Background()

	req := model.CreateContainerRequest{}

	err := c.Bind(&req)
	if err != nil {
		log.Panicf("binding body error: %s", err)
		return c.JSON(echo.ErrBadRequest.Code, nil)
	}

	err = h.srv.CreateContainer(ctx, req)
	if err != nil {
		return c.JSON(echo.ErrInternalServerError.Code, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
