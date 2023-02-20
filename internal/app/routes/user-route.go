package routes

import (
	"ticket-app/internal/app/controllers"
	"ticket-app/internal/app/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	version1 := "/api/v1"
	router.Use(middlewares.CORSMiddleware()).POST(version1+"/login", controllers.Login)
	v1 := router.Group(version1).Use(middlewares.Authentication(), middlewares.CORSMiddleware())
	{
		v1.POST("/ticket", controllers.CreateTicket)
		v1.POST("/ticket/search", controllers.ListTickets)
		v1.GET("/ticket/:id", controllers.GetTicket)
		v1.PUT("/ticket/:id", controllers.EditTicket)
		v1.DELETE("/ticket/:id", controllers.DeleteTicket)
		v1.PATCH("/ticket/change-status/:id", controllers.ChangeTicket)
		v1.PATCH("/ticket/change-priority/:id", controllers.ChangeTicket)
		v1.GET("/dashboard/count", controllers.TicketCount)
		v1.GET("/contacts", controllers.GetContacts)
		v1.GET("/users", controllers.GetUsers)
	}
}
