package main

import (
	"restapi-auth/controllers"
	"restapi-auth/database"
	"restapi-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
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
	api := router.Group("/api")
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
	return router
}
