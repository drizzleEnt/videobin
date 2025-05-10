package api

import (
	"fmt"
	"net/http"
	"videobin/internal/service"

	"github.com/gin-gonic/gin"
)

var _ Controller = (*handler)(nil)

type handler struct {
	service service.FileService
}

func New(s service.FileService) *handler {
	return &handler{
		service: s,
	}
}

// UploadFile implements Controller.
func (h *handler) UploadFile(c *gin.Context) {
	response := make(map[string]interface{}, 0)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response["error"] = fmt.Errorf("file field cannot be empty %w", err).Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		response["error"] = fmt.Errorf("file field open file %w", err).Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	filename := fileHeader.Filename

	if err := h.service.UploadFile(c.Request.Context(), filename, file); err != nil {
		response["error"] = fmt.Errorf("file cannot be uploaded %w", err).Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response["status"] = "ok"
	c.JSON(http.StatusOK, response)
}
