package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"todo-api-gin/libs/database"
)

type AuthHandlers struct{}

func (h *AuthHandlers) Register(c *gin.Context) {
	fmt.Println("Create")

	// get body from request and parse
	var body CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not possible to register: problem with password hashing"})
	}

	result := database.DB.Create(&User{
		Email:    body.Email,
		Password: string(hashedPass),
	})

	// return to user error in case it already exists
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			c.JSON(http.StatusConflict, gin.H{"error": "That email address already in use"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}

		return
	}

	// todo: create and encode password

	c.JSON(http.StatusCreated, nil)
}
