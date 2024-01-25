package main

import (
	"github.com/RajaSunrise/crud-fiber/database"
	"github.com/RajaSunrise/crud-fiber/database/migrations"
	"github.com/RajaSunrise/crud-fiber/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	migrations.Migrate()

	app := fiber.New()
	routers.SetRouters(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
