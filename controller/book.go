package controller

import (
	"go-jwt/database"
	"go-jwt/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type bookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func GetBooks(c *fiber.Ctx) error {
	var books []model.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		log.Printf("Error retrieving books: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var book model.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		log.Printf("Error retrieving book %d: %v", id, result.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var req bookRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Title == "" || req.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and author are required",
		})
	}

	book := model.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	result := database.DB.Create(&book)
	if result.Error != nil {
		log.Printf("Error creating book: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	log.Printf("Book created successfully: %v", book.Title)
	return c.Status(fiber.StatusCreated).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var req bookRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Title == "" || req.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and author are required",
		})
	}

	var book model.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		log.Printf("Error finding book %d: %v", id, result.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	book.Title = req.Title
	book.Author = req.Author

	result = database.DB.Save(&book)
	if result.Error != nil {
		log.Printf("Error updating book %d: %v", id, result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	log.Printf("Book updated successfully: %v", book.Title)
	return c.Status(fiber.StatusOK).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var book model.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		log.Printf("Error finding book %d: %v", id, result.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	result = database.DB.Delete(&book)
	if result.Error != nil {
		log.Printf("Error deleting book %d: %v", id, result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete book",
		})
	}

	log.Printf("Book deleted successfully: %d", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
