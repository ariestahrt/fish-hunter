package middlewares

import (
	"fish-hunter/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Admin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
		err := util.ValidateToken(tokenString)
		
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		
		return c.Next()
	}
}