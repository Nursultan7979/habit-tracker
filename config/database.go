package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"habit_tracker/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "user=postgres password=123456kz dbname=habit_tracker port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных: ", err)
		return
	}

	fmt.Println("Успешно подключено к базе данных")

	DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitLog{})
}
