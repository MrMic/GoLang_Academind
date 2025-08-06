package routes

import (
	"github.com/gin-gonic/gin"
	"michaelchlon.fr/api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // * NOTE: Register the handler for the "/events" route
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // * NOTE: Create a group for authenticated routes
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent     // * NOTE: Register the handler for the "/events/:id/register" route
		authenticated.DELETE("/events/:id/register", cancelRegistration) // * NOTE: Register the handler for the "/events/:id/register" route

	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup) // * NOTE: Register the handler for the "/signup" route
	server.POST("/login", login)   // * NOTE: Register the handler for the "/login" route
}
