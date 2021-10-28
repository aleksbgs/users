package main

import (
	"github.com/aleksbgs/users/src/database"
	"github.com/aleksbgs/users/src/routes"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":8001")
}



