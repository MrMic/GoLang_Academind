package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"michaelchlon.fr/api/models"
)

func signup(context *gin.Context) {
	var user models.User

	// * NOTE: Bind the JSON request body to the user instance
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.Save() // * NOTE: Save the user using the Save method defined in the model
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!", "user": user}) // * NOTE: Respond with a success message
}
