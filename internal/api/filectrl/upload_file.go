package filectrl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile implements Controller.
func (h *handler) UploadFile(c *gin.Context) {
	response := make(map[string]interface{}, 0)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response["error"] = fmt.Errorf("file field cannot be empty %w", err).Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.service.UploadFile(c.Request.Context(), fileHeader); err != nil {
		response["error"] = fmt.Errorf("file cannot be uploaded %w", err).Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response["status"] = "ok"
	c.JSON(http.StatusOK, response)
}
