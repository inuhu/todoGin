package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: "1", Task: "Learn GO"},
	{ID: "2", Task: "Wash the dish"},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
