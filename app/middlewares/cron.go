package middlewares

import (
	"fish-hunter/util"

	"github.com/gofiber/fiber/v2"
)

func Cron() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Print all header
		cronHeaderValue := c.GetReqHeaders()[util.GetConfig("CRON_HEADER_KEY")]

		// Compare
		if cronHeaderValue != util.GetConfig("CRON_HEADER_VALUE") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}