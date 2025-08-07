package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // * NOTE: Register the handler for the "/events" route
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup) // * NOTE: Register the handler for the "/signup" route
	server.POST("/login", login)   // * NOTE: Register the handler for the "/login" route
}
