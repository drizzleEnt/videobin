package api

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	UploadFile(c *gin.Context)
}
