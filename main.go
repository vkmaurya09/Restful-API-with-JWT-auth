package main

import (
	"restapi-auth/controllers"
	"restapi-auth/database"
	"restapi-auth/docs"
	"restapi-auth/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Task Management System"
	docs.SwaggerInfo.Description = "Crud API in Go Gin"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// connect database
	database.ConnectDatabase()
	// initilize all routes
	router := initRouter()
	// run server on 8080
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	// routes
	api := router.Group(docs.SwaggerInfo.BasePath)
	{
		// to register new user
		api.POST("/user/register", controllers.RegisterUser)
		// to get token for register users
		api.POST("/user/token", controllers.GenerateToken)
		// after succesful login you can use task manager api
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			// to get all tasks
			secured.GET("/tasks", controllers.AllTasks)
			// create new tasks
			secured.POST("/tasks", controllers.CreateTask)
			// find particular task
			secured.GET("/tasks/one", controllers.FindTask)
			// update any task by id
			secured.PUT("/tasks/update", controllers.UpdateTask)
			// delete any task by id
			secured.DELETE("/tasks/delete", controllers.DeleteTask)

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
