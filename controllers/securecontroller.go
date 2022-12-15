package controllers

import (
	"fmt"
	"net/http"
	"restapi-auth/database"
	"restapi-auth/models"
	"time"

	"github.com/gin-gonic/gin"
)

// create task struct
type CreateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

// update task struct
type UpdateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

// get all task
func AllTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// create task
func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentTime := time.Now()

	// Create task
	task := models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail, Date: currentTime.Format("2006.01.02 15:04:05")}
	database.DB.Create(&task)

	// return created task
	c.JSON(http.StatusOK, gin.H{"data": task})

}

// find task particular task by id
func FindTask(c *gin.Context) {
	var task models.Task
	fmt.Println(c.Request.URL.Query())
	id := c.Request.URL.Query().Get("id")
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}

// update task by id
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// delete Task
func DeleteTask(c *gin.Context) {
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
