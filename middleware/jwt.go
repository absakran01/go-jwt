package middleware

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error { 
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("Unauthorized access attempt or error: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})(c)
}