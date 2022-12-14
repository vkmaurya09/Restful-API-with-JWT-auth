package models

type Task struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}
