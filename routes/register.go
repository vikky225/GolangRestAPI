package routes

import (
	"net/http"
	"strconv"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"could not fetch error from DB": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in register for event": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registered for event", "event": event, "user": userId})
}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in cancel registration": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration cancelled for event", "event": event, "user": userId})
}
