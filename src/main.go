package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/Routes"
	"github.com/ienjir/ArtaferaBackend/src/api/artPicture"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/database/sampledata"
	plsthisisjustappackge "github.com/ienjir/ArtaferaBackend/src/minio"
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

	err = plsthisisjustappackge.InitMinIO()
	if err != nil {
		return
	}

	err = sampledata.SeedDatabase()
	if err != nil {
		return
	}

	router.Static("/files", artPicture.UploadDir)

	Routes.RegisterRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
