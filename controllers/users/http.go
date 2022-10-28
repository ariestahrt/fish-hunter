package users

import (
	"fish-hunter/businesses/users"
	"fish-hunter/controllers/users/requests"
	"fish-hunter/controllers/users/response"
	"fish-hunter/helpers"
	"fish-hunter/util"
	"fmt"
	"strings"

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
	userInput := requests.UserRegister{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": helpers.StrUnprocessableEntity,
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	user, err := ctrl.authUseCase.Register(userInput.ToDomain())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println(user)

	return c.Status(fiber.StatusCreated).JSON(response.FromDomain(user))
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	userInput := requests.UserLogin{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": helpers.StrUnprocessableEntity,
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := ctrl.authUseCase.Login(userInput.ToDomain())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})	
}

func (ctrl *AuthController) GetProfile(c *fiber.Ctx) error {
	// Get id from JWT
	tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
	JWTClaim := util.GetJWTPayload(tokenString)

	user, err := ctrl.authUseCase.GetProfile(JWTClaim.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.FromDomain(user))
}

func (ctrl *AuthController) UpdateProfile(c *fiber.Ctx) error {
	// Get ID From JWT

	return nil
}