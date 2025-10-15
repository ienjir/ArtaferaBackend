package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/database/sampledata"
	miniobucket "github.com/ienjir/ArtaferaBackend/src/minio"
	"github.com/ienjir/ArtaferaBackend/src/routes"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("test")
	// Load env's from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	utils.SetGinMode()

	router := gin.Default()
	router.Use(utils.CORSMiddleware())

	auth.LoadAuthEnvs()
	validation.LoadsValidationEnvs()

	err = auth.GenerateNewArgon2idHash()
	if err != nil {
		return
	}

	err = router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		return
	}

	err = database.ConnectDatabase()
	if err != nil {
		return
	}

	err = miniobucket.InitMinIO()
	if err != nil {
		return
	}

	if utils.GinMode != 2 {
		log.Println("Deleting all buckets")
		if err := miniobucket.DeleteAllBuckets(); err != nil {
			return
		}
	}

	if err := miniobucket.CreateMinioBuckets(); err != nil {
		return
	}

	if utils.GinMode != 2 {
		if err := sampledata.SeedDatabase(); err != nil {
			fmt.Printf("Error while seeding database: %s", err.Error())
		}
	}

	routes.RegisterRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
