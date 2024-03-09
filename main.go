package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itoodua12/Fiber-Rest-API/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to api")
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
