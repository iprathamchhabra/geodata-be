package config

import (
	"fmt"
	"geo-data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "host=dpg-ctv2mk1u0jms73as2iqg-a user=goe_data_database_user password=BTY0odnVwmhBFsCrJ2J6H0bKbiw1vUqz dbname=goe_data_database port=5432 TimeZone=Asia/Shanghai"
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
