package controller

import (
	"go-jwt/database"
	"go-jwt/model"
	"go-jwt/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *fiber.Ctx) error {
	user, err := parseAuthRequest(ctx)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	log.Printf("User registered successfully: %v", user.Email)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func Login(ctx *fiber.Ctx) error {
	user, err := parseAuthRequest(ctx)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var existingUser model.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		log.Printf("Error finding user: %v", result.Error)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to find user, check email",
		})
	}

	if !utils.CheckPass(user.Password, existingUser.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token, err := utils.GenerateToken(existingUser.ID)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	log.Printf("token generated for user: %v", existingUser)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func parseAuthRequest(ctx *fiber.Ctx) (model.User, error) {
	var authRequest authRequest
	err := ctx.BodyParser(&authRequest)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Email: authRequest.Email,
		Password: utils.HashPass(authRequest.Password),
	}
	return user, nil
}