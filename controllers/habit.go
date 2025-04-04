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
	config.DB.Create(&habit)
	c.JSON(http.StatusOK, habit)
}

func GetHabits(c *gin.Context) {
	var habits []models.Habit
	if err := config.DB.Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

func UpdateHabit(c *gin.Context) {
	var habit models.Habit
	if err := config.DB.Where("id = ?", c.Param("id")).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена"})
		return
	}

	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&habit)
	c.JSON(http.StatusOK, habit)
}

func DeleteHabit(c *gin.Context) {
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&models.Habit{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Привычка не найдена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Привычка удалена"})
}
