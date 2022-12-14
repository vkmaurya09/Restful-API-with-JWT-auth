package controllers

import (
	"fmt"
	"net/http"
	"restapi-auth/database"
	"restapi-auth/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

type CreateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

type UpdateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

// all task will show
func AllTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// create task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentTime := time.Now()

	// Create task
	task := models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail, Date: currentTime.Format("2006.01.02 15:04:05")}
	database.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

// find task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task
	fmt.Println(c.Request.URL.Query())
	id := c.Request.URL.Query().Get("id")
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteUser(c *gin.Context) {
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	database.DB.First(&task, id)
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		database.DB.Delete(&task)
		c.JSON(200, gin.H{"success": "Task#" + id + " deleted"})
	}
}
