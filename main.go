package main

import (
	"fmt"
	"net/http"

	"example.com/restapi/db"
	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/event", getEvent)
	server.POST("/event", createEvent)

	server.Run(":8080") // localhost : 8080

	fmt.Println("Hello, World!")

}

func getEvent(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error in get event": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
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
