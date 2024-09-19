package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"todo-api-gin/libs/database"
)

type AuthHandlers struct{}

func (h *AuthHandlers) Register(c *gin.Context) {
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

	c.JSON(http.StatusCreated, nil)
}

func (h *AuthHandlers) Login(c *gin.Context) {
	// parse json
	var body LoginUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo:  get user from db by checking if password is correct
	// todo: generate gwt token
}
