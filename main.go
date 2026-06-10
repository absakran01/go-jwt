package main

import (
	"go-jwt/database"
	"go-jwt/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.InitDB()

	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
