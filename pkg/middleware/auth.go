package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtTokenContextKey = "user" // i.e. default
	jwtCustomAdminKey  = "admin"
)

func OnlyAdmin(c *fiber.Ctx) error {
	user := c.Locals(jwtTokenContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims[jwtCustomAdminKey].(bool)
	if !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}
	// Go to next middleware:
	return c.Next()
}
