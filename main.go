package main

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/config"
	"habit_tracker/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
