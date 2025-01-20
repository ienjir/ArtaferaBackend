package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/Routes"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/database/sampledata"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()

	// Load env's from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Set vars in the packages
	auth.LoadAuthEnvs()
	validation.LoadsValidationEnvs()

	// Generate Argon2idHash for password hashing
	err = auth.GenerateNewArgon2idHash()
	if err != nil {
		return
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
	err = sampledata.GenerateFakeData()
	if err != nil {
		return
	}

	err2 := auth.VerifyToken("xeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG4uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzM3NDYwNjk2LCJpZCI6MSwicm9sZSI6IkFkbWluIn0.RG1eh5PNr4Byehuz7m2UA-4xj0G3AVnYDJyxkKboBbc")
	if err2 != nil {
		log.Fatal(err2.Message)
	}

	// Register routes
	Routes.RegisterRoutes(router, database.DB)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
