package handlers

import (
	"modules/modules/config"
	"modules/modules/entities"

	"github.com/gofiber/fiber/v2"
)

func AddMyCars(c *fiber.Ctx) error {
	cars := new(entities.CarProducts)

	if err := c.BodyParser(cars); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.DataBase.Create(&cars)
	return c.Status(200).JSON(cars)
}
