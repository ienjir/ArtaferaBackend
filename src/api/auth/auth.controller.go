package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

var Argon2IDHash *Argon2idHash
var MinEntropyBits float64
var JWTSecret string

func HashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)

	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}

	return hashSalt, nil
}

func LoginHandler(c *gin.Context) {
	var json models.LoginRequest

	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
