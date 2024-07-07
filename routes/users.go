package routes

import (
	"net/http"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in signup": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user created": user})
}
