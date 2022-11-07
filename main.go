package main

import (
	"fish-hunter/app/routes"
	drivers "fish-hunter/drivers"
	mongo_driver "fish-hunter/drivers/mongo"
	"fish-hunter/util"
	"fish-hunter/util/datasetutil"
	"fish-hunter/util/s3"
	"fish-hunter/util/scrapper"
	"fmt"

	_userUseCase "fish-hunter/businesses/users"
	_userController "fish-hunter/controllers/users"

	_urlUseCase "fish-hunter/businesses/urls"
	_urlController "fish-hunter/controllers/urls"

	_jobUseCase "fish-hunter/businesses/jobs"
	_jobController "fish-hunter/controllers/jobs"

	_datasetUseCase "fish-hunter/businesses/datasets"
	_datasetController "fish-hunter/controllers/datasets"

	_statUseCase "fish-hunter/businesses/stats"
	_statController "fish-hunter/controllers/stats"

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
	urlScrapper := scrapper.NewUrlScrapper()
	urlRepo := drivers.NewUrlRepository(mongo_driver.GetDB())
	urlUsecase := _urlUseCase.NewUrlUseCase(urlRepo, urlScrapper)
	urlController := _urlController.NewUrlController(urlUsecase)

	// Job
	jobRepo := drivers.NewJobRepository(mongo_driver.GetDB())
	jobUsecase := _jobUseCase.NewJobUseCase(jobRepo)
	jobController := _jobController.NewJobController(jobUsecase)

	// S3
	esTiga := s3.NewAWS_S3()

	// Dataset
	datasetUtil := datasetutil.NewDatasetUtil()
	datasetRepo := drivers.NewDatasetRepository(mongo_driver.GetDB())
	datasetUsecase := _datasetUseCase.NewDatasetUseCase(datasetRepo, esTiga, datasetUtil)
	datasetController := _datasetController.NewDatasetController(datasetUsecase, userUsecase)

	// Stat
	statUsecase := _statUseCase.NewStatUseCase(datasetRepo, jobRepo, urlRepo)
	statController := _statController.NewStatController(statUsecase)

	// Setup Routes
	routes := routes.ControllerList{
		UserController: *userController,
		UrlController: *urlController,
		JobController: *jobController,
		DatasetController: *datasetController,
		StatController: *statController,
	}
	
	routes.Setup(app)
	
	app.ListenTLS(util.GetConfig("APP_PORT"), util.GetConfig("TLS_CERT"), util.GetConfig("TLS_KEY"))
}