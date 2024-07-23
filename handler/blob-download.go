package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Siravitt/azure-storage/model"
	"github.com/labstack/echo/v4"
)

func (h handler) BlobDownload(c echo.Context) error {
	ctx := context.Background()

	req := model.BlobDownloadRequest{}

	err := c.Bind(&req)
	if err != nil {
		log.Panicf("binding body error: %s", err)
		return c.JSON(echo.ErrBadRequest.Code, nil)
	}

	resp, err := h.srv.BlobDownload(ctx, req)
	if err != nil {
		return c.JSON(echo.ErrInternalServerError.Code, nil)
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename="+req.FileName)
	c.Response().Header().Set("Content-Type", "application/octet-stream")

	return c.HTMLBlob(http.StatusOK, resp.Bytes())
}
