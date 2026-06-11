package middleware

import (
	"go-jwt/database"
	"go-jwt/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthMiddleware() fiber.Handler {
	
	return basicauth.New(basicauth.Config{
		
		Authorizer: func(username, password string) bool {
			dbUser := model.User{}
			res := database.DB.Where("email = ?", username).First(&dbUser)
			if res.Error != nil {
				log.Printf("Error finding user: %v", res.Error)
				return false
			}
			
			err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
			if err != nil {
				log.Printf("Error comparing password: %v", err)
				return false
			}
			return true
		},

	})
}
