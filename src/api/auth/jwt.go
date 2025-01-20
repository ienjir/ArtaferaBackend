package auth

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"os"
	"time"
)

var JWTSecret []byte

func LoadAuthEnvs() {
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(User models.User) (*jwt2.Token, error) {
	// Create a new token object
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"email": User.Email,
		"id":    User.ID,
		"role":  User.Role.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	/* tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	} */

	return token, nil
}
