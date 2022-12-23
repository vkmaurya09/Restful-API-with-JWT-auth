package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi-auth/database"
	"restapi-auth/models"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// pATH Information

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

// AllTasks godoc
// @Security bearerAuth
// @Summary Show all tasks
// @Description Get all tasks
// @Accept  json
// @Produce json
// @Success 200 {array} models.Task
// @Router  /secured/tasks [get]
func AllTasks(c *gin.Context) {
	// connection to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var tasks []models.Task
	val, err := client.Get("tasks").Result()
	if err == nil {
		json.Unmarshal([]byte(val), &tasks)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"data from redis": tasks})
	} else {
		database.DB.Find(&tasks)
		json, err := json.Marshal(tasks)
		if err != nil {
			fmt.Println(err)
		}

		err = client.Set("tasks", json, time.Minute).Err()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"data from db": tasks})
		return
	}

	// fmt.Println(val)

}

// CreateTask godoc
// @Security bearerAuth
// @Summary Create new tasks
// @Description create tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskInput true "Create task"
// @Success 200 {object} models.Task
// @Router  /secured/tasks [post]t]
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

// FindTask godoc
// @Security bearerAuth
// @Summary Find Single Task
// @Description Find single task
// @Accept json
// @Produce json
// @Param  id query int true "task Id"
// @Success 200 {object} models.Task
// @Router  /secured/tasks/one [get]
func FindTask(c *gin.Context) {
	// connection to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// task := models.Task{}
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	val, err := client.Get(id).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &id)
		if err != nil {
			panic(err)
		}
		// return from redis
		c.JSON(http.StatusOK, gin.H{"data from redis": val})
	} else {
		if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		// marshal json
		json, err := json.Marshal(task)
		if err != nil {
			fmt.Println(err)
		}
		// save json by id in redis
		err = client.Set("id", json, time.Minute).Err()
		if err != nil {
			fmt.Println(err)
		}
		// return from db
		c.JSON(http.StatusOK, gin.H{"data from db": task})
		return
	}
}

// UpdateTask godoc
// @Security bearerAuth
// @Summary Update task
// @Description Update tasks
// @Accept json
// @Produce json
// @Param  id query int true "task Id"
// @Param task body UpdateTaskInput true "Update task"
// @Success 200 {object} models.Task
// @Router  /secured/tasks/update [put]
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

// DeleteTask godoc
// @Security bearerAuth
// @Summary delete task
// @Description delete task
// @Accept json
// @Produce json
// @Param  id query int true "task Id"
// @Success 200 {object} models.Task
// @Router  /secured/tasks/delete [delete]
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	database.DB.First(&task, id)
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&task)
	c.JSON(200, gin.H{"success": "Task " + id + " deleted"})

}
