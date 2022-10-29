package main

import (
	"fish-hunter/app/routes"
	drivers "fish-hunter/drivers"
	mongo_driver "fish-hunter/drivers/mongo"
	"fish-hunter/util"
	"fmt"

	_userUseCase "fish-hunter/businesses/users"
	_userController "fish-hunter/controllers/users"

	_cronUseCase "fish-hunter/businesses/cron"
	_cronController "fish-hunter/controllers/cron"

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

	// Cron
	cronRepo := drivers.NewCronRepository(mongo_driver.GetDB())
	cronUsecase := _cronUseCase.NewCronUseCase(cronRepo)
	cronController := _cronController.NewCronController(cronUsecase)

	// Setup Routes
	routes := routes.ControllerList{
		UserController: *userController,
		CronController: *cronController,
	}
	routes.Setup(app)
	
	app.ListenTLS(util.GetConfig("APP_PORT"), util.GetConfig("TLS_CERT"), util.GetConfig("TLS_KEY"))
}