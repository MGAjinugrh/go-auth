package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("[user]:[password]/[dbname]"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to DB")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
