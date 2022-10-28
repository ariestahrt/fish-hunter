package routes

import (
	"fish-hunter/app/middlewares"
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
	group.Get("/stats", middlewares.Roles([]string{"admin", "guest"}), func (c *fiber.Ctx) error {
		return c.SendString("HELLO STATS")
	})

	group.Get("/top_brands", nil)
	group.Get("/stats_by_day", nil)

	// Jobs
	group.Get("/jobs", nil)

	// Urls
	group.Get("/urls", nil)

	// Users
	group.Get("/user", middlewares.Authorized(), cl.UserController.GetProfile)
	group.Post("/user/register", cl.UserController.Register)
	group.Post("/user/login", cl.UserController.Login)
	group.Put("/user/update_profile", nil)
	group.Put("/user/update_password", nil)

	group.Get("/logout", nil)
	group.Get("/validate", nil)
}