package users

import (
	"fish-hunter/businesses/users"
	"fish-hunter/controllers/users/requests"
	"fish-hunter/controllers/users/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUseCase users.UseCase
}

func NewAuthController(authUseCase users.UseCase) *AuthController {
	return &AuthController{
		authUseCase,
	}
}

func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	userInput := requests.User{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println(userInput.ToDomain())
	
	user := ctrl.authUseCase.Register(userInput.ToDomain())
	return c.Status(fiber.StatusCreated).JSON(response.FromDomain(user))
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	userInput := requests.UserLogin{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user,err := ctrl.authUseCase.Login(userInput.ToDomain())
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(response.FromDomain(user))
}
