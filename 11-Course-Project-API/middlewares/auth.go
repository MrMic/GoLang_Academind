package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"michaelchlon.fr/api/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: No token provided.",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid token.",
		})
		return
	}

	// NOTE: 📢 Store the user ID in the context for later use
	context.Set("userId", userId)

	context.Next()
}
