package auth

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"os"
	"strconv"
)

// HashSalt struct used to store generated hash and salt used to generate the hash.
type HashSalt struct {
	Hash, Salt []byte
}

type Argon2idHash struct {
	time    uint32 // time represents the number of passed over the specified memory.
	memory  uint32 // cpu memory to be used.
	threads uint8  // threads for parallelism aspect of the algorithm.
	keyLen  uint32 // keyLen of the generate hash key.
	saltLen uint32 // saltLen the length of the salt used.
}

// NewArgon2idHash constructor function for Argon2idHash.
func NewArgon2idHash(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

// GenerateHash using the password and provided salt. If not salt value provided fallback to random value generated of a given length.
func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error

	// If salt is not provided generate a salt of the configured salt length.
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}

	if err != nil {
		return nil, err
	}

	// Generate hash
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)

	// Return the generated hash and salt used for storage.
	return &HashSalt{Hash: hash, Salt: salt}, nil
}

// Compare generated hash with store hash.
func (a *Argon2idHash) Compare(hash, salt, password []byte) error {

	// Generate hash for comparison.
	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return err
	}

	// Compare the generated hash with the stored hash. If they don't match return error.
	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("hash doesn't match")
	}

	return nil
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func GenerateNewArgon2idHash() error {
	hashTimeStr := os.Getenv("HASH_TIME")
	hashSaltLengthStr := os.Getenv("HASH_SALT_LENGTH")
	hashMemoryStr := os.Getenv("HASH_MEMORY")
	hashThreadsStr := os.Getenv("HASH_THREADS")
	hashKeyLengthStr := os.Getenv("HASH_KEY_LENGTH")

	// Convert to appropriate types
	hashTime, err := strconv.ParseUint(hashTimeStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_TIME: %v\n", err)
		return err
	}

	hashSaltLength, err := strconv.ParseUint(hashSaltLengthStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_SALT_LENGTH: %v\n", err)
		return err
	}

	hashMemory, err := strconv.ParseUint(hashMemoryStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_MEMORY: %v\n", err)
		return err
	}

	hashThreads, err := strconv.ParseUint(hashThreadsStr, 10, 8)
	if err != nil {
		fmt.Printf("Error converting HASH_THREADS: %v\n", err)
		return err
	}

	hashKeyLength, err := strconv.ParseUint(hashKeyLengthStr, 10, 32)
	if err != nil {
		fmt.Printf("Error converting HASH_KEY_LENGTH: %v\n", err)
		return err
	}

	// Pass converted values to the function
	Argon2IDHash = NewArgon2idHash(
		uint32(hashTime),
		uint32(hashSaltLength),
		uint32(hashMemory),
		uint8(hashThreads),
		uint32(hashKeyLength),
	)

	return nil
}
