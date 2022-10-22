package main

import (
	"final-project-1/config"
	"final-project-1/httpserver/controllers"
	"final-project-1/httpserver/repositories"
	"final-project-1/httpserver/routers"
	"final-project-1/httpserver/services"
	"log"

	_ "final-project-1/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3030
// @BasePath /api
func main() {

	err := godotenv.Load(".env.dev")

	if err != nil {
		log.Fatal("Environment Variables not found")
	}

	app := gin.Default()
	appRoute := app.Group("/api")
	db, _ := config.Connect()

	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	routers.TodoRouter(appRoute, todoController)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":3030")
}
