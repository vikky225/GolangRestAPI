// RegisteredRoutes initializes routes for the gin server
package routes

import (
	"example.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine) {
	// Define the routes and their handlers
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// We can protect route like this in Golang using middlewares
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//server.POST("/events", middlewares.Authenticate, createEvents)
	//server.PUT("/events/:id", updateEvent)
	//server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
