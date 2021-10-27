package main

import (
	"github.com/aleksbgs/users/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()


	app.Listen(":8001")
}



