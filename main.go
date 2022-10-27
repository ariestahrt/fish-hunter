package main

import (
	"fish-hunter/app/routes"
	drivers "fish-hunter/drivers"
	mongo_driver "fish-hunter/drivers/mongo"
	"fish-hunter/util"
	"fmt"

	_userUseCase "fish-hunter/businesses/users"
	_userController "fish-hunter/controllers/users"

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

	userRepo := drivers.NewUserRepository(mongo_driver.GetDB())
	userUsecase := _userUseCase.NewUserUseCase(userRepo)
	userController := _userController.NewAuthController(userUsecase)
	// Setup Routes
	routes := routes.ControllerList{
		UserController: *userController,
	}
	routes.Setup(app)
	
	app.Listen(util.GetConfig("APP_PORT"))
}