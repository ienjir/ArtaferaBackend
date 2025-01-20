package auth

import (
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"log"
	"net/http"
	"os"
	"time"
)

var JWTSecret []byte

func LoadAuthEnvs() {
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(JWTSecret) == 0 {
		log.Fatal("JWT_SECRET environment variable is required")
	}
}

func GenerateJWT(User models.User) (*string, *models.ServiceError) {
	// Create a new token object
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"email": User.Email,
		"id":    User.ID,
		"role":  User.Role.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to sign JWT token"}
	}

	return &tokenString, nil
}

func VerifyToken(tokenString string) *models.ServiceError {
	token, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			// Returning ServiceError if signing method is unexpected
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: fmt.Sprintf("Error verifying token: %v", err)}
	}

	if !token.Valid {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "Token is invalid"}
	}

	return nil
}
