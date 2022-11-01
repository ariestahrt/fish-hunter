package middlewares

import (
	"encoding/base64"
	"encoding/json"
	appjwt "fish-hunter/util/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Roles(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
		err := appjwt.ValidateToken(tokenString)
		
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Check for admin role
		payload := strings.Split(tokenString, ".")[1]
		e := base64.StdEncoding.WithPadding(base64.NoPadding)
		decoded, _ := e.DecodeString(payload)
		decodedString := string(decoded)
		
		payloadJSON := appjwt.JWTClaim{}

		json.Unmarshal([]byte(decodedString), &payloadJSON)

		for _, role := range payloadJSON.Roles {
			for _, allowedRole := range roles {
				if role == allowedRole {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Role Unauthorized",
		})
	}
}