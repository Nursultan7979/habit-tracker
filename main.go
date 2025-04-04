package main

import (
	"github.com/gin-gonic/gin"
	"habit_tracker/config"
	"habit_tracker/routes"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	config.ConnectDatabase()

	r.Run(":8080")
}
