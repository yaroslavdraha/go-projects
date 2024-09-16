package todos

import (
	"time"
	"todo-api-gin/libs/database"
)

type Todo struct {
	database.Model
	Title       string     `json:"title"`
	Status      string     `json:"status"`
	CompletedAt *time.Time `json:"completed_at"`
}

type CreateTodo struct {
	Title string `json:"title" binding:"required"`
}
