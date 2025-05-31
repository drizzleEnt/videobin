package api

import (
	"github.com/gin-gonic/gin"
)

type FileController interface {
	UploadFile(c *gin.Context)
}
