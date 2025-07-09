package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // * NOTE: Register the handler for the "/events" route

	server.Run(":8082") // * NOTE: Start the server on port 8082 (localhost:8082)
}

// * NOTE: Define a handler for the "/events" route -----------------------------
func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"}) // * NOTE: Respond with a JSON object containing a message
}
