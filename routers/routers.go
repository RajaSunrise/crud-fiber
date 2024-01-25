package routers

import (
	"github.com/RajaSunrise/crud-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/users/:id", controllers.GetUserByID)
	app.Post("/api/users", controllers.CreateUser)
	app.Put("/api/users/:id", controllers.UpdateUser) // Fix: Include the ":id" parameter in the route
	app.Delete("/api/users/:id", controllers.DeleteUser)
}
