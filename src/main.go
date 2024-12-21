package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/routes"
)

func main() {
	router := gin.Default()

	// Initialize the database
	database.ConnectDatabase()

	// If not existing make the file storage directory
	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server started")
}
