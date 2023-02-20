package main

import (
	"ticket-app/internal/app/routes"
	database "ticket-app/internal/app/utils/database"
	"ticket-app/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetConfigParams()
	router := gin.Default()
	routes.UserRoute(router)
	database.ConnectDB(config.Config.MongoDB)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	router.Run()
}
