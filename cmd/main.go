package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/database"
	"github.com/ienjir/ArtaferaBackend/routes"
)

func main() {
	router := gin.Default()

	// Initialize the database
	database.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server started")
}
