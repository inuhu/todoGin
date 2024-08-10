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

func RenderTodos(c *gin.Context) {
	var todos []Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"todos": todos,
	})
}

// func GetTodos(c *gin.Context) {
// 	var todos []Todo
// 	if err := db.DB.Find(&todos).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, todos)

// }

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Если произошла другая ошибка, возвращаем 500 статус с сообщением об ошибке
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		return
	}
	c.JSON(http.StatusOK, todo)

}

func CreateTodo(c *gin.Context) {
	var newTodo Todo
	newTodo.Task = c.PostForm("task")
	newTodo.Description = c.PostForm("description")
	db.DB.Create(&newTodo)
	c.Redirect(http.StatusFound, "/todos")
}

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
