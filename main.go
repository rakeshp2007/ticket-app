package main

import (
	"ticket-app/configuration"
	"ticket-app/routes"
	"ticket-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration.SetConfigParams()
	router := gin.Default()
	routes.UserRoute(router)
	services.ConnectDB(configuration.Config.MongoDB)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	//token, _ := services.GenerateJWT("rakesh.raghavan@aspiresys.com")
	//fmt.Println(token)
	//val := services.ValidateToken(token)
	router.Run()
}
