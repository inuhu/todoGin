package todo

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inuhu/todo-app/pkg/db"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
}

func GetTodos(c *gin.Context) {
	var todos []Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)

}

func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}
	db.DB.Create(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// func DeleteTodo(c *gin.Context) {
// 	id := c.Param("id")
// 	var todo Todo
// 	if err := db.DB.First(&todo, id); err != nil {
// 		c.Status(http.StatusNotFound)
// 		return
// 	}
// 	db.DB.Delete(&todo)
// 	c.Status(http.StatusNoContent)
// }

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo

	// Проверяем, существует ли запись
	if err := db.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Удаляем запись и проверяем на ошибки
	if err := db.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
