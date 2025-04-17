package main

import (
	"tgn-backend/connection"
	"tgn-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	connection.CreateConnection()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}
