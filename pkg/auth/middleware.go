package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT token from Authorization header
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authorization token",
		})
	}

	// Remove "Bearer " prefix
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	claims, err := ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token: " + err.Error(),
		})
	}

	// Store claims in context for later use
	c.Locals("user_id", claims.UserID)
	c.Locals("email", claims.Email)
	c.Locals("roles", claims.Roles)

	return c.Next()
}
