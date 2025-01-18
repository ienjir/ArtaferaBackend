package auth

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"golang.org/x/crypto/argon2"
	"net/http"
	"os"
	"strconv"
	"time"
)

var JWTSecret []byte
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

func LoadAuthEnvs() {
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

func VerifyUser(request models.LoginRequest) (*models.User, *models.ServiceError) {
	var User models.User

	// Check if user exists
	if err := database.DB.Where("email = ?", request.Email).First(&User).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
	}

	if err := ComparePassword(User, request.Password); err != nil {
		return nil, &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	return &User, nil
}

func ComparePassword(user models.User, password string) *models.ServiceError {
	var Password, Hash, Salt []byte

	Password = []byte(password)
	Hash = user.Password
	Salt = user.Salt

	err := Argon2IDHash.Compare(Hash, Salt, Password)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "Password is wrong"}
	}

	return nil
}

func GenerateJWT(User models.User) (string, error) {
	// Create a new token object
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"email": User.Email,
		"id":    User.ID,
		"role": ,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JW)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
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

/*
	Validation
*/

func VerifyLoginData(Data models.LoginRequest) *models.ServiceError {
	err := validation.ValidatePassword(Data.Password)
	if err != nil {
		return err
	}

	err = validation.ValidateEmail(Data.Email)
	if err != nil {
		return err
	}

	return nil
}
