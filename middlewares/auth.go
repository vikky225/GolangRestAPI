package middlewares

import (
	"net/http"

	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// we should extract token if we want to protect route for Signed in User JWT
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"No unauthrieed ": err.Error()})
		return
	}
	context.Set("userId", userId)

	context.Next()
}
