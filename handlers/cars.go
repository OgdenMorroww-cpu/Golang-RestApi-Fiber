package handlers

import (
	"modules/modules/config"
	"modules/modules/entities"

	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {
	var cars []entities.CarProducts

	config.DataBase.Find(&cars)
	return c.Status(200).JSON(cars)
}
