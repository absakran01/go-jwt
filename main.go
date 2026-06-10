package main

import (
	"go-jwt/controller"
	"go-jwt/database"
	"go-jwt/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header missing",
		})
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	valid, err := utils.ValidateToken(parts[1])
	if err != nil || !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	return c.Next()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.InitDB()

	app := fiber.New()

	auth := app.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)

	books := app.Group("/books", jwtMiddleware)
	books.Get("/", controller.GetBooks)
	books.Get("/:id", controller.GetBook)
	books.Post("/", controller.CreateBook)
	books.Put("/:id", controller.UpdateBook)
	books.Delete("/:id", controller.DeleteBook)

	log.Fatal(app.Listen(":3000"))
}