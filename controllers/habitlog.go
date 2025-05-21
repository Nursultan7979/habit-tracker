package controllers

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/config"
	"habit_tracker/models"
	"net/http"
	"strconv"
)

func CreateHabitLog(c *gin.Context) {
	var input models.HabitLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	input.UserID = userID.(uint)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании лога привычки"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func GetHabitLogs(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var logs []models.HabitLog
	if err := config.DB.Where("user_id = ?", userID).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении логов"})
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

	id := c.Param("id")
	var habitLog models.HabitLog
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&habitLog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден или не принадлежит вам"})
		return
	}

	var input struct {
		HabitID uint   `json:"habit_id"`
		Date    string `json:"date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var habit models.Habit
	if err := config.DB.Where("id = ? AND user_id = ?", input.HabitID, userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Эта привычка вам не принадлежит"})
		return
	}

	habitLog.Date = input.Date

	if err := config.DB.Save(&habitLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении лога"})
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

	id := c.Param("id")
	result := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.HabitLog{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Лог не найден или не принадлежит вам"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Лог удален"})
}

//Admin

func GetAllHabitLogs(c *gin.Context) {
	var logs []models.HabitLog
	if err := config.DB.Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить журналы привычек"})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func UpdateHabitLogByAdmin(c *gin.Context) {
	id := c.Param("id")
	var log models.HabitLog
	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Журнал привычки не найден"})
		return
	}

	var input struct {
		HabitID uint `json:"habit_id"`
		UserID  uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.HabitID != 0 {
		log.HabitID = input.HabitID
	}
	if input.UserID != 0 {
		log.UserID = input.UserID
	}

	if err := config.DB.Save(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Журнал привычки обновлён"})
}

func DeleteHabitLogByAdmin(c *gin.Context) {
	id := c.Param("id")

	logID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	result := config.DB.Delete(&models.HabitLog{}, logID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Журнал привычки не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Журнал привычки удалён"})
}
