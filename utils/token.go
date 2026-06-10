package utils

func GenerateToken(userID uint) (string, error) {
	jwtSecret := []byte("your_secret_key") // In production, use an environment variable
}