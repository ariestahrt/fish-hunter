package middlewares

import (
	appjwt "fish-hunter/util/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
		err := appjwt.ValidateToken(tokenString)
		
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}