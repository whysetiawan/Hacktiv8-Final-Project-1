package routers

import (
	controllers "final-project-1/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRouter(route *gin.Engine, todoController controllers.TodoController) {
	todoRouter := route.Group("/todo")
	{
		todoRouter.POST("", todoController.CreateTodo)
		todoRouter.DELETE(":id", todoController.DeleteTodo)
	}
}
