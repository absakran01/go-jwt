package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func GenerateToken(userID uint) (string, error) {
	if jwtSecret == nil {
		jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	if jwtSecret == nil {
		jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return false, err
	}
	return token.Valid, nil
}