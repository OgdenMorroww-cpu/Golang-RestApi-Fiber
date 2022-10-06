package servers

import (
	"log"
	"modules/modules/config"
	"modules/modules/handlers"

	"github.com/gofiber/fiber/v2"
)

func Servers() {
	app := fiber.New()

	config.ConnectDataBase()

	app.Get("/cars", handlers.GetCars)
	app.Get("/cars/:id", handlers.GetMyCar)
	app.Post("/cars", handlers.AddMyCars)
	app.Put("/cars/:id", handlers.UpdateMyCars)
	app.Delete("/cars/:id", handlers.RemoveMyCars)

	log.Fatal(app.Listen(":5000"))
}
