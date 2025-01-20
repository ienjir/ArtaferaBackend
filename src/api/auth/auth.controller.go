package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

var Argon2IDHash *Argon2idHash

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
	}

	if err := VerifyLoginData(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	user, err := VerifyUser(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	jwt, err2 := GenerateJWT(*user)
	if err2 != nil {
		return
	}

	fmt.Println(jwt)
}
