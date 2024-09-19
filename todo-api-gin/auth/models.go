package auth

import (
	"todo-api-gin/libs/database"
)

type User struct {
	database.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string
}

type CreateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
