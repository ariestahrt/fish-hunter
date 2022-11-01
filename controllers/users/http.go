package users

import (
	"fish-hunter/businesses/users"
	"fish-hunter/controllers/users/requests"
	"fish-hunter/controllers/users/response"
	"fish-hunter/helpers"
	appjwt "fish-hunter/util/jwt"
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
	JWTClaim := appjwt.GetJWTPayload(tokenString)

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
	tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
	JWTClaim := appjwt.GetJWTPayload(tokenString)

	userInput := requests.UserUpdateProfile{}

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

	_, err := ctrl.authUseCase.UpdateProfile(userInput.ToDomain(JWTClaim.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update profile success",
	})
}

func (ctrl *AuthController) UpdatePassword(c *fiber.Ctx) error {
	// Get ID From JWT
	tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
	JWTClaim := appjwt.GetJWTPayload(tokenString)

	userInput := requests.UserUpdatePassword{}

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

	_, err := ctrl.authUseCase.UpdatePassword(userInput.ToDomain(JWTClaim.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update password success",
	})
}

func (ctrl *AuthController) GetAllUsers(c *fiber.Ctx) error {
	users, err := ctrl.authUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.FromDomainArray(users))
}

func (ctrl *AuthController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.authUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.FromDomain(user))
}

func (ctrl *AuthController) UpdateByAdmin(c *fiber.Ctx) error {
	id := c.Params("id")

	userInput := requests.UserUpdateByAdmin{}

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

	_, err := ctrl.authUseCase.Update(userInput.ToDomain(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update user success",
	})
}

func (ctrl *AuthController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := ctrl.authUseCase.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete user success",
	})
}

func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	// Get ID From JWT
	tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)

	ctrl.authUseCase.Logout(tokenString)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout success",
	})
}