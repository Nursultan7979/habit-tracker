package routes

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.POST("/habits", controllers.CreateHabit)
	router.GET("/habits", controllers.GetHabits)
	router.PUT("/habits/:id", controllers.UpdateHabit)
	router.DELETE("/habits/:id", controllers.DeleteHabit)

	router.POST("/habitlogs", controllers.CreateHabitLog)
	router.GET("/habitlogs", controllers.GetHabitLogs)
	router.PUT("/habitlogs/:id", controllers.UpdateHabitLog)
	router.DELETE("/habitlogs/:id", controllers.DeleteHabitLog)
}
