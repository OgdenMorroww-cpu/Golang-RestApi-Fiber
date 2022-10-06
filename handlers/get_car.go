package handlers

import (
	"modules/modules/config"
	"modules/modules/entities"

	"github.com/gofiber/fiber/v2"
)

func GetMyCar(c *fiber.Ctx) error {
	id := c.Params("id")
	var cars entities.CarProducts

	result := config.DataBase.Find(&cars, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&cars)
}
