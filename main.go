package main

import (
	"fmt"

	"example.com/restapi/db"
	"example.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisteredRoutes(server)

	server.Run(":8080") // localhost : 8080

	fmt.Println("Hello, World!")

}
