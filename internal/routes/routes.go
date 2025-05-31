package routes

import (
	"fmt"
	"videobin/internal/api"
	"videobin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	c api.FileController,
	mw middleware.Middleware,
) *gin.Engine {
	r := gin.New()

	api := r.Group("/files")
	api.Use(mw.Check)
	api.POST("/upload", c.UploadFile)
	api.POST("/download", func(ctx *gin.Context) { fmt.Println("download") })

	auth := r.Group("/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) { fmt.Println("login hello") })
		auth.POST("/register", func(ctx *gin.Context) { fmt.Println("register hello") })
	}

	return r
}
