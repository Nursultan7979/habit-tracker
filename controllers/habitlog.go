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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	habitLog.UserID = userID.(uint)

	if err := config.DB.Create(&habitLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habitLog)
}

func GetHabitLogs(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var logs []models.HabitLog
	if err := config.DB.Where("user_id = ?", userID).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func UpdateHabitLog(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var habitLog models.HabitLog
	if err := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&habitLog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден или не принадлежит вам"})
		return
	}

	if err := c.ShouldBindJSON(&habitLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&habitLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habitLog)
}

func DeleteHabitLog(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	if err := config.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).Delete(&models.HabitLog{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден или не принадлежит вам"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Лог удален"})
}
