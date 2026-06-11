package router

import (
	"go-jwt/controller"
	"go-jwt/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)

	books := api.Group("/books")
	books.Get("/", controller.GetBooks)
	books.Get("/:id", controller.GetBook)

	//secured routes
	books.Use(middleware.BasicAuthMiddleware())
	books.Post("/", controller.CreateBook)
	books.Put("/:id", controller.UpdateBook)
	books.Delete("/:id", controller.DeleteBook)
}
