package database

import (
	"log"
	"restapi-auth/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB instance
var DB *gorm.DB

// db connection function
func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:Vkm@12345@tcp(127.0.0.1:3306)/restfulapi?parseTime=true")

	if err != nil {
		log.Fatal(err)
		panic("Cannot Connect to DB")
	}
	log.Println("Connected to DB")

	// automigrate USers
	database.AutoMigrate(&models.User{})
	// automograte tasks
	database.AutoMigrate(&models.Task{})

	DB = database
}
