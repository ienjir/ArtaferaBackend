package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/Routes"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set proxies
	err = router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		return
	}

	// Initialize the database
	err = database.ConnectDatabase()
	if err != nil {
		return
	}

	// Generate fake data to
	database.GenerateFakeData(database.DB)

	// Register routes
	Routes.RegisterRoutes(router, database.DB)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
