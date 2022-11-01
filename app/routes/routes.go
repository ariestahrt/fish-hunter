package routes

import (
	"fish-hunter/app/middlewares"
	"fish-hunter/controllers/datasets"
	"fish-hunter/controllers/jobs"
	"fish-hunter/controllers/stats"
	"fish-hunter/controllers/urls"
	"fish-hunter/controllers/users"

	"github.com/gofiber/fiber/v2"
)

type ControllerList struct {
	UserController users.AuthController
	UrlController  urls.UrlController
	JobController  jobs.JobController
	DatasetController datasets.DatasetController
	StatController stats.StatsController
}

func (cl *ControllerList) Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	group := app.Group("/api/v1")

	group.Get("/stats/all", cl.StatController.GetStatistics)
	group.Get("/stats/lastweek", middlewares.Authorized(), cl.StatController.GetLastWeekStatistics)

	// Datasets
	group.Get("/datasets/top_brands", middlewares.Authorized(), cl.DatasetController.TopBrands)
	group.Get("/datasets/status/:status", middlewares.Authorized(), cl.DatasetController.Status)
	group.Get("/datasets/view/:id/*", cl.DatasetController.View)
	group.Get("/datasets/:id/activate", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Activate)
	group.Put("/datasets/:id/validate", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Validate)
	group.Get("/datasets/:id/download", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Download)
	group.Get("/datasets/:id", middlewares.Authorized(), cl.DatasetController.GetByID)

	// Jobs
	group.Get("/jobs", middlewares.Authorized(), cl.JobController.GetAll)
	group.Get("/jobs/:id", middlewares.Authorized(), cl.JobController.GetByID)

	// Urls
	group.Get("/urls", middlewares.Authorized(), cl.UrlController.GetAll)
	group.Get("/urls/:id", middlewares.Authorized(), cl.UrlController.GetByID)
	group.Get("/urls/fetch/:source", middlewares.Cron(), cl.UrlController.FetchUrl)

	// Users
	group.Get("/user", middlewares.Authorized(), cl.UserController.GetProfile)
	group.Post("/user/register", cl.UserController.Register)
	group.Post("/user/login", cl.UserController.Login)
	group.Put("/user/update_profile", middlewares.Authorized(), cl.UserController.UpdateProfile)
	group.Put("/user/update_password", middlewares.Authorized(), cl.UserController.UpdatePassword)
	group.Get("/user/all", middlewares.Roles([]string{"admin"}), cl.UserController.GetAllUsers)
	group.Get("/user/logout", middlewares.Authorized(), cl.UserController.Logout)
	group.Get("/user/:id", middlewares.Roles([]string{"admin"}), cl.UserController.GetByID)
	group.Put("/user/:id", middlewares.Roles([]string{"admin"}), cl.UserController.UpdateByAdmin)
	group.Delete("/user/:id", middlewares.Roles([]string{"admin"}), cl.UserController.Delete)
}