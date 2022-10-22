package main

import (
	"final-project-1/config"
	"final-project-1/docs"
	"final-project-1/httpserver/controllers"
	"final-project-1/httpserver/repositories"
	"final-project-1/httpserver/routers"
	"final-project-1/httpserver/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

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

	docs.SwaggerInfo.Title = "Hacktiv8 Final-Project-1 API"
	docs.SwaggerInfo.Description = "This is just a simple TODO List"
	docs.SwaggerInfo.Host = "hacktiv8-final-project-1-production.up.railway.app"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run()
}
