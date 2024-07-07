package routes

import (
	"net/http"
	"strconv"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}

	e, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in get event by id": err.Error()})
		return
	}
	context.JSON(http.StatusOK, e)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in get event": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {

	//Authenticate(context)

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userId")
	err := models.Save(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in create event": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event created": event}) // event)
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"error in get event by id": err.Error()})
		return
	}

	var updatedEvent models.Event

	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEvent.ID = eventId
	err = models.Update(updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in update event": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event updated": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error parsing": err.Error()})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in get event by id": err.Error()})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in delete event": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event deleted": id})
}
