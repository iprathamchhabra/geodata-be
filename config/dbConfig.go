package config

import (
	"fmt"
	"geo-data/models"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%s", host, user, dbname, sslmode, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.GeoData{}, &models.User{}, &models.Shape{})
	if err != nil {
		fmt.Println("Failed to auto migrate models:", err)
		return nil, err
	}

	fmt.Println("Auto migration successful.")
	
	return db, nil
}