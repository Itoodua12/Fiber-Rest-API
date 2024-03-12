package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itoodua12/Fiber-Rest-API/database"
	"github.com/itoodua12/Fiber-Rest-API/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
	// Products Endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

}

func main() {
	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
