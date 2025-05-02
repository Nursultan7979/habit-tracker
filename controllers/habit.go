package controllers

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/config"
	"habit_tracker/models"
	"net/http"
)

func CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	habit.UserID = userID.(uint)
	if err := config.DB.Create(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habit)
}

func GetHabits(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var habits []models.Habit
	if err := config.DB.Where("user_id = ?", userID).Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habits)
}

func UpdateHabit(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var habit models.Habit
	if err := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена или не принадлежит вам"})
		return
	}

	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habit)
}

func DeleteHabit(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	if err := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).Delete(&models.Habit{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена или не принадлежит вам"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Привычка удалена"})
}

// admin
func GetAllHabits(c *gin.Context) {
	var habits []models.Habit
	if err := config.DB.Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить привычки"})
		return
	}
	c.JSON(http.StatusOK, habits)
}

func UpdateHabitByAdmin(c *gin.Context) {
	id := c.Param("id")
	var habit models.Habit
	if err := config.DB.First(&habit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена"})
		return
	}

	var input models.Habit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit.Title = input.Title
	habit.Description = input.Description

	if err := config.DB.Save(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Привычка обновлена"})
}

func DeleteHabitByAdmin(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Habit{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Привычка удалена"})
}
