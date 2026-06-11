package middleware

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error { 
	return jwtware.New(jwtware.Config{
		TokenLookup: "cookie:jwt",
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("JWT validation error: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT token",
			})
		},
	})(c)
}