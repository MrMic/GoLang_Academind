package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"michaelchlon.fr/api/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId") // * NOTE: Get the user ID from the context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err := event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for the event successfully!"})
}

func cancelRegistration(context *gin.Context) {}
