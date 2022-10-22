package main

import (
	"final-project-1/config"
	"final-project-1/httpserver/controllers"
	"final-project-1/httpserver/repositories"
	"final-project-1/httpserver/routers"
	"final-project-1/httpserver/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env.dev")

	if err != nil {
		log.Fatal("Environment Variables not found")
	}

	app := gin.Default()
	db, _ := config.Connect()

	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	routers.TodoRouter(app, todoController)

	app.Run(":3030")
}
