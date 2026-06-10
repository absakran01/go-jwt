package database

import (
	"fmt"
	"log"
	"os"
	"go-jwt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=go_jwt sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Database connection established")

	// Auto migrate the User and Book models
	err = DB.AutoMigrate(&model.User{}, &model.Book{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}
	log.Println("Database migrated successfully")
}