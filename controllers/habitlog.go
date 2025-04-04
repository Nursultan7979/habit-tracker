package controllers

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/config"
	"habit_tracker/models"
	"net/http"
)

func CreateHabitLog(c *gin.Context) {
	var habitLog models.HabitLog
	if err := c.ShouldBindJSON(&habitLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&habitLog)
	c.JSON(http.StatusOK, habitLog)
}

func GetHabitLogs(c *gin.Context) {
	var logs []models.HabitLog
	if err := config.DB.Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func UpdateHabitLog(c *gin.Context) {
	var habitLog models.HabitLog
	if err := config.DB.Where("id = ?", c.Param("id")).First(&habitLog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден"})
		return
	}

	if err := c.ShouldBindJSON(&habitLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&habitLog)
	c.JSON(http.StatusOK, habitLog)
}

func DeleteHabitLog(c *gin.Context) {
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&models.HabitLog{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Лог удален"})
}
