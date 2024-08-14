package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuhu/todo-app/pkg/db"
	"github.com/inuhu/todo-app/pkg/todo"
)

func main() {
	db.Init()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/todos", todo.RenderTodos)
	r.GET("/todo/:id", todo.GetTodo)
	r.POST("/todos", todo.CreateTodo)
	r.POST("/delete/:id", todo.DeleteTodo)
	r.GET("/edit/:id", todo.EditTodoForm)
	r.POST("/update/:id", todo.UpdateTodo)
	r.Run(":8080")
}
