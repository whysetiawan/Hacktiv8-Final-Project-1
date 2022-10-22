package controllers

import (
	"final-project-1/httpserver/dto"
	"final-project-1/httpserver/services"
	"net/http"

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

func (c *todoController) CreateTodo(ctx *gin.Context) {
	var dto dto.CreateTodoDto
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	todo, err := c.todoService.CreateTodo(dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error")
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *todoController) DeleteTodo(ctx *gin.Context) {}
