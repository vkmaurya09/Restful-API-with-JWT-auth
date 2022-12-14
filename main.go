package main

import (
	"restapi-auth/controllers"
	"restapi-auth/database"
	"restapi-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	// database.Connect("root:Vkm@12345@tcp(127.0.0.1:3306)/restfulapi?parseTime=true")
	// database.Migrate()
	database.ConnectDatabase()
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/tasks", controllers.AllTasks)
			secured.POST("/tasks", controllers.CreateTask)
			secured.GET("/tasks/one", controllers.FindTask)
			secured.PUT("/tasks/update", controllers.UpdateTask)
			secured.DELETE("/tasks/delete", controllers.DeleteUser)

		}
	}
	return router
}
