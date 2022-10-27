package routes

import (
	"fish-hunter/controllers/users"

	"github.com/gofiber/fiber/v2"
)

type ControllerList struct {
	UserController users.AuthController
}

func (cl *ControllerList) Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	group := app.Group("/api/v1")
	group.Get("/stats", nil)
	group.Get("/top_brands", nil)
	group.Get("/stats_by_day", nil)

	// Jobs
	group.Get("/jobs", nil)

	// Urls
	group.Get("/urls", nil)

	// Users
	group.Post("/register", cl.UserController.Register)
	group.Post("/login", cl.UserController.Login)
	group.Get("/logout", nil)
	group.Get("/validate", nil)
}