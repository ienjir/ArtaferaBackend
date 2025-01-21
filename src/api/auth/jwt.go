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

var (
	JWTAccessSecret  []byte
	JWTRefreshSecret []byte
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoadAuthEnvs() {
	JWTAccessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	JWTRefreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

	if len(JWTAccessSecret) == 0 || len(JWTRefreshSecret) == 0 {
		log.Fatal("JWT_ACCESS_SECRET and JWT_REFRESH_SECRET environment variables are required")
	}
}

func GenerateTokenPair(user models.User) (*TokenPair, *models.ServiceError) {
	// Generate access token
	accessToken := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"role":  user.Role.Role,
		"type":  "access",
		"exp":   time.Now().Add(30 * time.Minute).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(JWTAccessSecret)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to generate access token",
		}
	}

	// Generate refresh token
	refreshToken := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"id":   user.ID,
		"type": "refresh",
		"exp":  time.Now().Add(168 * time.Hour).Unix(), // 7 days
	})

	refreshTokenString, err := refreshToken.SignedString(JWTRefreshSecret)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to generate refresh token",
		}
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func VerifyAccessToken(tokenString string) (*jwt2.Token, *models.ServiceError) {
	return verifyToken(tokenString, JWTAccessSecret, "access")
}

func VerifyRefreshToken(tokenString string) (*jwt2.Token, *models.ServiceError) {
	return verifyToken(tokenString, JWTRefreshSecret, "refresh")
}

func verifyToken(tokenString string, secret []byte, tokenType string) (*jwt2.Token, *models.ServiceError) {
	token, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    fmt.Sprintf("Invalid token: %v", err),
		}
	}

	claims, ok := token.Claims.(jwt2.MapClaims)
	if !ok || !token.Valid {
		return nil, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token claims",
		}
	}

	// Verify token type
	if claimType, ok := claims["type"].(string); !ok || claimType != tokenType {
		return nil, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token type",
		}
	}

	return token, nil
}

func RefreshTokens(refreshToken string) (*TokenPair, *models.ServiceError) {
	// Verify the refresh token
	token, err := VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt2.MapClaims)
	userID := claims["id"]

	// Here you would typically look up the user in your database
	// For this example, we'll create a minimal user object
	user := models.User{
		Firstname: userID.(string),
		// You should populate other fields from your database
	}

	// Generate new token pair
	return GenerateTokenPair(user)
}
