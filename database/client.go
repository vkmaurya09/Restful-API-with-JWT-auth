package database

import (
	"log"
	"restapi-auth/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "gorm.io/gorm"
)

// var Instance *gorm.DB

// func Connect(connectionString string) {
// 	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
// 	if dbError != nil {
// 		log.Fatal(dbError)
// 		panic("Cannot Connect to DB")
// 	}
// 	log.Println("Connected to DB")
// }

// func Migrate() {
// 	Instance.AutoMigrate(&models.User{})
// 	Instance.AutoMigrate(&models.Task{})
// 	log.Println("DB Migration COmpleted")
// }

var DB *gorm.DB

// var dbError error

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:Vkm@12345@tcp(127.0.0.1:3306)/restfulapi?parseTime=true")

	if err != nil {
		log.Fatal(err)
		panic("Cannot Connect to DB")
	}
	log.Println("Connected to DB")
	database.AutoMigrate(&models.User{})

	database.AutoMigrate(&models.Task{})

	DB = database
}
