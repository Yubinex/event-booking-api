package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yubinex/event-booking-api/db"
	"github.com/yubinex/event-booking-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
