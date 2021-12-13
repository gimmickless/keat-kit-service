package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	ctxTokenKey        = "user"
	claimAdminCheckKey = "admin"
)

func isAdmin(c *fiber.Ctx) bool {
	token := c.Locals(ctxTokenKey).(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims[claimAdminCheckKey].(bool)
}
