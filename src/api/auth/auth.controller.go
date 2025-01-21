package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

var Argon2IDHash Argon2idHash

func HashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)
	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}
	return hashSalt, nil
}

func Login(c *gin.Context) {
	var json models.LoginRequest
	if jsonErr := c.ShouldBindJSON(&json); jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonErr.Error()})
		return
	}

	if err := VerifyLoginData(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	user, err := VerifyUser(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	jwt, err2 := GenerateTokenPair(*user)
	if err2 != nil {
		c.JSON(err2.StatusCode, gin.H{"error": err2.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})
}

func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.GetHeader("X-Refresh-Token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Refresh token is required",
		})
		return
	}

	newTokens, err := RefreshTokens(refreshToken)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, newTokens)
}
