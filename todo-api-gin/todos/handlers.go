package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-api-gin/libs/database"
)

type TodoHandlers struct{}

func (h *TodoHandlers) Create(c *gin.Context) {
	var body CreateTodo

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&Todo{
		Title: body.Title,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *TodoHandlers) List(c *gin.Context) {
	var todos []Todo
	result := database.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandlers) Delete(c *gin.Context) {
	result := database.DB.Delete(&Todo{}, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
