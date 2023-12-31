package main

import (
	"fiber/database"
	"fiber/database/migration"
	"fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init fiber
	app := fiber.New()
	// Init database
	database.DatabaseInit()
	migration.Migration()

	// route
	route.RouteInit(app)

	app.Listen(":8080")
}
