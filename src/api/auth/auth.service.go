package auth

import (
	"fmt"
	"os"
	"strconv"
)

var Argon2IDHash *Argon2idHash

func hashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)

	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}

	return hashSalt, nil
}

func GenerateNewArgon2idHash() {
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
	Argon2IDHash = NewArgon2idHash(
		uint32(hashTime),
		uint32(hashSaltLength),
		uint32(hashMemory),
		uint8(hashThreads),
		uint32(hashKeyLength),
	)
}
