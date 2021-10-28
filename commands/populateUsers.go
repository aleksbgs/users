package main

import (
	"github.com/aleksbgs/users/src/database"
	"github.com/aleksbgs/users/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	database.Connect()
	dsn := "host=" + "host.docker.internal" + " " + "user=" + "root" + " " + "password=" + "root" + " " + "dbname=" + "ambassador" + " " + "port=5432 sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var users []models.User
	DB.Find(&users)

	for _, user := range users {
		database.DB.Create(&user)
	}

}
