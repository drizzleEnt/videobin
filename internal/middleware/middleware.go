package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func New() Middleware {
	return &middlewareController{}
}

type middlewareController struct {
}

// Check implements Middleware.
func (m *middlewareController) Check(c *gin.Context) {
	fmt.Println("checking")
}
