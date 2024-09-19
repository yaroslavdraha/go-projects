package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo-api-gin/auth"
	"todo-api-gin/libs/database"
	"todo-api-gin/todos"
)

var todosHandler = todos.TodoHandlers{}
var authHandler = auth.AuthHandlers{}

func main() {
	// Connect to database
	database.Connect("storage/sqlite.db")
	database.Migrate(&todos.Todo{}, &auth.User{})

	r := gin.Default()

	// Public routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	r.Use(authMiddleware())

	// Protected routes
	todosRoutes := r.Group("/todos")
	{
		todosRoutes.POST("", todosHandler.Create)
		todosRoutes.GET("", todosHandler.List)
		todosRoutes.DELETE("/:id", todosHandler.Delete)
	}

	usersRoutes := r.Group("/users")
	{
		usersRoutes.POST("/current", authHandler.Register)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Auth middleware")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
