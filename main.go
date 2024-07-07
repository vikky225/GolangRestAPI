package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/restapi/db"
	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080") // localhost : 8080

	fmt.Println("Hello, World!")

}

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
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//event.ID = 1
	//event.UserID = 1
	err := models.Save(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in create event": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event created": event}) // event)
}
