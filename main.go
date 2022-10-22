package main

import (
	"final-project-1/httpserver/controllers"
	"final-project-1/httpserver/routers"
	"final-project-1/httpserver/services"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	todoService := services.NewTodoService()
	todoController := controllers.NewTodoController(todoService)

	routers.TodoRouter(app, todoController)

	app.Run(":3030")
}
