package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var GinMode int

func SetGinMode() {
	mode := os.Getenv("MODE")

	// Default to release in production environments
	if mode == "" {
		mode = "release"
		GinMode = 2
		log.Println("MODE not set, defaulting to 'release'")
	}

	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
		GinMode = 0
		fmt.Println("Running in debug mode")
	case "test":
		gin.SetMode(gin.TestMode)
		GinMode = 1
		fmt.Println("Running in test mode")
	case "release":
		gin.SetMode(gin.ReleaseMode)
		GinMode = 2
		fmt.Println("Running in release mode")
	default:
		log.Fatalf("Invalid MODE '%s'. Must be one of: debug, test, release", mode)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // Cache preflight for 24 hours

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func GenerateRandomFileName() string {
	rand.Seed(time.Now().UnixNano())
	data := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000000))

	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
