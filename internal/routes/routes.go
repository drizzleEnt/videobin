package routes

import (
	"videobin/internal/api"

	"github.com/gin-gonic/gin"
)

func InitRoutes(c api.Controller) *gin.Engine {
	r := gin.New()

	api := r.Group("/files")
	{
		api.POST("/upload", c.UploadFile)
	}

	return r
}
