package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/Routes"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var argon2IDHash *auth.Argon2idHash

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	generateNewArgon2idHash()

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

func generateNewArgon2idHash() {
	hashTimeStr := os.Getenv("HASH_TIME")
	hashSaltLengthStr := os.Getenv("HASH_SALT_LENGTH")
	hashMemoryStr := os.Getenv("HASH_MEMORY")
	hashThreadsStr := os.Getenv("HASH_THREADS")
	hashKeyLengthStr := os.Getenv("HASH_KEY_LENGTH")

	// Convert to appropriate types
	hashTime, err := strconv.ParseUint(hashTimeStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_TIME: %v\n", err)
		return
	}

	hashSaltLength, err := strconv.ParseUint(hashSaltLengthStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_SALT_LENGTH: %v\n", err)
		return
	}

	hashMemory, err := strconv.ParseUint(hashMemoryStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_MEMORY: %v\n", err)
		return
	}

	hashThreads, err := strconv.ParseUint(hashThreadsStr, 10, 8)
	if err != nil {
		fmt.Printf("Error converting HASH_THREADS: %v\n", err)
		return
	}

	hashKeyLength, err := strconv.ParseUint(hashKeyLengthStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_KEY_LENGTH: %v\n", err)
		return
	}

	// Pass converted values to the function
	argon2IDHash = auth.NewArgon2idHash(
		uint32(hashTime),
		uint32(hashSaltLength),
		uint32(hashMemory),
		uint8(hashThreads),
		uint32(hashKeyLength),
	)
}
