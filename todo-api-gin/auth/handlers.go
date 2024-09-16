package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthHandlers struct{}

func (h *AuthHandlers) Register(c *gin.Context) {
	fmt.Println("Create")
}
