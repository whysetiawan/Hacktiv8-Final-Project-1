package routers

import (
	controllers "final-project-1/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRouter(route *gin.RouterGroup, todoController controllers.TodoController) *gin.RouterGroup {
	todoRouter := route.Group("/todo")
	{
		todoRouter.GET("", todoController.GetAllTodos)
		todoRouter.GET(":id", todoController.GetTodoByID)
		todoRouter.POST("", todoController.CreateTodo)
		todoRouter.DELETE(":id", todoController.DeleteTodo)
		todoRouter.PUT("/update/:id", todoController.UpdateTodo)
	}
	return todoRouter
}
