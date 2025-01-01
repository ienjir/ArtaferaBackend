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

	// Generate fake data to
	database.GenerateFakeData(database.DB)

	// Register routes
	routes.RegisterRoutes(router, database.DB)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server started")
}
