// RegisteredRoutes initializes routes for the gin server
package routes

import "github.com/gin-gonic/gin"

func RegisteredRoutes(server *gin.Engine) {
	// Define the routes and their handlers
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
}
