package middleware

import (
	"pade-backend/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware wraps the auth package middleware
func AuthMiddleware(c *fiber.Ctx) error {
	return auth.AuthMiddleware(c)
}

// RoleMiddleware checks if user has one of the required roles
func RoleMiddleware(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, ok := c.Locals("roles").([]string)
		if !ok || len(roles) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		hasRole := false
		for _, userRole := range roles {
			for _, required := range requiredRoles {
				if userRole == required {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "insufficient permissions",
			})
		}

		return c.Next()
	}
}
