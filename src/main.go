package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/routes"
)

func main() {
	router := gin.Default()

	// Set proxies
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		return
	}

	// Initialize the database
	database.ConnectDatabase()

	// Generate fake data to
	database.GenerateFakeData(database.DB)

	// Register routes
	routes.RegisterRoutes(router, database.DB)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
