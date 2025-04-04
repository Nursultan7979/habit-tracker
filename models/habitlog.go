package models

import "gorm.io/gorm"

type HabitLog struct {
	gorm.Model
	HabitID uint `json:"habit_id"`
	Habit   Habit
	Date    string `json:"date"` // формат: YYYY-MM-DD
}
