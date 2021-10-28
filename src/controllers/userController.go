package controllers

import (
	"github.com/aleksbgs/users/src/database"
	"github.com/aleksbgs/users/src/models"
	"github.com/gofiber/fiber"
	"strconv"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	IsAmbassador, _ := strconv.ParseBool(data["is_ambassador"])

	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: IsAmbassador,
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}
