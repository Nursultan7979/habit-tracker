package config

import (
	"fmt"
	"os"
	"time"

	"habit_tracker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Retrying to connect to DB in 3 seconds...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitLog{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
