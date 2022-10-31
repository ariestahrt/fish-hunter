package main

import (
	"fish-hunter/app/routes"
	drivers "fish-hunter/drivers"
	mongo_driver "fish-hunter/drivers/mongo"
	"fish-hunter/util"
	"fmt"

	_userUseCase "fish-hunter/businesses/users"
	_userController "fish-hunter/controllers/users"

	_urlUseCase "fish-hunter/businesses/urls"
	_urlController "fish-hunter/controllers/urls"

	_jobUseCase "fish-hunter/businesses/jobs"
	_jobController "fish-hunter/controllers/jobs"

	_datasetUseCase "fish-hunter/businesses/datasets"
	_datasetController "fish-hunter/controllers/datasets"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Starting server...")

	app := fiber.New()
	app.Use(cors.New())

	// Setup DB
	client, err := mongo_driver.Connect()
	if err != nil {
		panic(err)
	}
	mongo_driver.SetClient(client)

	// User
	userRepo := drivers.NewUserRepository(mongo_driver.GetDB())
	userUsecase := _userUseCase.NewUserUseCase(userRepo)
	userController := _userController.NewAuthController(userUsecase)

	// Url
	urlRepo := drivers.NewUrlRepository(mongo_driver.GetDB())
	urlUsecase := _urlUseCase.NewUrlUseCase(urlRepo)
	urlController := _urlController.NewUrlController(urlUsecase)

	// Job
	jobRepo := drivers.NewJobRepository(mongo_driver.GetDB())
	jobUsecase := _jobUseCase.NewJobUseCase(jobRepo)
	jobController := _jobController.NewJobController(jobUsecase)

	// Dataset
	datasetRepo := drivers.NewDatasetRepository(mongo_driver.GetDB())
	datasetUsecase := _datasetUseCase.NewDatasetUseCase(datasetRepo)
	datasetController := _datasetController.NewDatasetController(datasetUsecase)

	// Setup Routes
	routes := routes.ControllerList{
		UserController: *userController,
		UrlController: *urlController,
		JobController: *jobController,
		DatasetController: *datasetController,
	}
	
	routes.Setup(app)
	
	app.ListenTLS(util.GetConfig("APP_PORT"), util.GetConfig("TLS_CERT"), util.GetConfig("TLS_KEY"))
}