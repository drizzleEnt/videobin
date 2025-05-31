package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	Check(c *gin.Context)
}
