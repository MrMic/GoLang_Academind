package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // * NOTE: Register the handler for the "/events" route
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
}
