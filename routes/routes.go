package routes

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/controllers"
	"habit_tracker/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware)

	auth.GET("/profile", controllers.GetProfile)
	auth.PUT("/profile", controllers.UpdateUser)
	auth.DELETE("/profile", controllers.DeleteUser)

	auth.POST("/habits", controllers.CreateHabit)
	auth.GET("/habits", controllers.GetHabits)
	auth.PUT("/habits/:id", controllers.UpdateHabit)
	auth.DELETE("/habits/:id", controllers.DeleteHabit)

	auth.POST("/habit-logs", controllers.CreateHabitLog)
	auth.GET("/habit-logs", controllers.GetHabitLogs)
	auth.PUT("/habit-logs/:id", controllers.UpdateHabitLog)
	auth.DELETE("/habit-logs/:id", controllers.DeleteHabitLog)

	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware)
	admin.Use(middleware.RequireRole("admin"))

	admin.GET("/users", controllers.GetUsers)
	admin.PUT("/users/:id", controllers.UpdateUserByAdmin)
	admin.DELETE("/users/:id", controllers.DeleteUserByAdmin)
	admin.POST("/users", controllers.CreateUserByAdmin)

	admin.GET("/habits", controllers.GetAllHabits)
	admin.PUT("/habits/:id", controllers.UpdateHabitByAdmin)
	admin.DELETE("/habits/:id", controllers.DeleteHabitByAdmin)

	admin.GET("/habit-logs", controllers.GetAllHabitLogs)
	admin.PUT("/habit-logs/:id", controllers.UpdateHabitLogByAdmin)
	admin.DELETE("/habit-logs/:id", controllers.DeleteHabitLogByAdmin)
}
