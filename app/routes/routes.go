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

	api := app.Group("/api/v1")
	stats := api.Group("/stats")
	
	stats.Get("/all", cl.StatController.GetStatistics)
	stats.Get("/lastweek", middlewares.Authorized(), cl.StatController.GetLastWeekStatistics)

	// Datasets
	datasets := api.Group("/datasets")
	datasets.Get("/top_brands", middlewares.Authorized(), cl.DatasetController.TopBrands)
	datasets.Get("/status/:status", middlewares.Authorized(), cl.DatasetController.Status)
	datasets.Get("/view/:id/*", cl.DatasetController.View)
	datasets.Get("/:id/activate", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Activate)
	datasets.Put("/:id/validate", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Validate)
	datasets.Get("/:id/download", middlewares.Roles([]string{"admin", "user"}), cl.DatasetController.Download)
	datasets.Get("/:id", middlewares.Authorized(), cl.DatasetController.GetByID)

	// Jobs
	jobs := api.Group("/jobs")
	jobs.Get("", middlewares.Authorized(), cl.JobController.GetAll)
	jobs.Get("/:id", middlewares.Authorized(), cl.JobController.GetByID)

	// Urls
	urls := api.Group("/urls")
	urls.Get("", middlewares.Authorized(), cl.UrlController.GetAll)
	urls.Get("/:id", middlewares.Authorized(), cl.UrlController.GetByID)
	urls.Get("/fetch/:source", middlewares.Cron(), cl.UrlController.FetchUrl)

	// Users
	user := api.Group("/user")
	user.Get("", middlewares.Authorized(), cl.UserController.GetProfile)
	user.Post("/register", cl.UserController.Register)
	user.Post("/login", cl.UserController.Login)
	user.Put("/update_profile", middlewares.Authorized(), cl.UserController.UpdateProfile)
	user.Put("/update_password", middlewares.Authorized(), cl.UserController.UpdatePassword)
	user.Get("/all", middlewares.Roles([]string{"admin"}), cl.UserController.GetAllUsers)
	user.Get("/:id", middlewares.Roles([]string{"admin"}), cl.UserController.GetByID)
	user.Put("/:id", middlewares.Roles([]string{"admin"}), cl.UserController.UpdateByAdmin)
	user.Delete("/:id", middlewares.Roles([]string{"admin"}), cl.UserController.Delete)
}