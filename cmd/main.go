package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuhu/todo-app/pkg/todo"
)

func main() {
	r := gin.Default()

	r.GET("/todos", todo.GetTodos)
	r.POST("/todos", todo.CreateTodo)
	r.DELETE("/todos/:id", todo.DeleteTodo)
	r.Run(":9999")
}
