package main

import (
	"github.com/gin-gonic/gin"
	"todo-api-gin/auth"
	"todo-api-gin/libs/database"
	"todo-api-gin/todos"
)

var todosHandler = todos.TodoHandlers{}
var authHandler = auth.AuthHandlers{}

func main() {
	r := gin.Default()

	// Connect to database
	database.Connect("storage/sqlite.db")
	database.Migrate(&todos.Todo{}, &auth.User{})

	// Define routes
	todosRoutes := r.Group("/todos")
	{
		todosRoutes.POST("/", todosHandler.Create)
		todosRoutes.GET("/", todosHandler.List)
		todosRoutes.DELETE("/:id", todosHandler.Delete)
	}

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Register)
		authRoutes.POST("/me", authHandler.Register)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
