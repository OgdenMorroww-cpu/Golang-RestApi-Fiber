package handlers

import (
	"modules/modules/config"
	"modules/modules/entities"

	"github.com/gofiber/fiber/v2"
)

func RemoveMyCars(c *fiber.Ctx) error {
	id := c.Params("id")
	var cars entities.CarProducts

	result := config.DataBase.Delete(&cars, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
