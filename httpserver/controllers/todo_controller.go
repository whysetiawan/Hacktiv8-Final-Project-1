package controllers

import (
	"final-project-1/httpserver/services"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) *todoController {
	return &todoController{
		todoService,
	}
}

func (c *todoController) CreateTodo(ctx *gin.Context) {}

func (c *todoController) DeleteTodo(ctx *gin.Context) {}
