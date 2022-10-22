package controllers

import (
	"final-project-1/httpserver/dto"
	"final-project-1/httpserver/services"
	"final-project-1/utils"
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

// CreateTodo godoc
// @Tags    Todo
// @Summary create a todo
// @Param   todo body     dto.CreateTodoDto true "Create Todo DTO"
// @Success 200  {object} utils.HttpSuccess[models.TodoModel]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /todo [post]
func (c *todoController) CreateTodo(ctx *gin.Context) {
	var dto dto.CreateTodoDto
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	todo, err := c.todoService.CreateTodo(dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Todo Created", todo))
}

// DeleteTodo godoc
// @Tags    Todo
// @Summary create a todo
// @Param   id  path     int true "Todo ID"
// @Success 200 {object} utils.HttpSuccess[models.TodoModel]
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router  /todo [delete]
func (c *todoController) DeleteTodo(ctx *gin.Context) {}
