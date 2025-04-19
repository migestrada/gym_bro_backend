package main

import (
	"gym-bro-backend/connection"
	"gym-bro-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	connection.CreateConnection()
	var router *gin.Engine = gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
