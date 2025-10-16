package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strings"
)

var Argon2IDHash Argon2idHash

func Login(c *gin.Context) {
	var json models.LoginRequest
	if jsonErr := c.ShouldBindJSON(&json); jsonErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.Email = strings.ToLower(json.Email)

	if err := VerifyLoginData(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := VerifyUser(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	jwt, err2 := GenerateTokenPair(*user)
	if err2 != nil {
		utils.RespondWithServiceError(c, err2)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"token": jwt})
}

func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.GetHeader("X-Refresh-Token")
	if refreshToken == "" {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrRefreshTokenRequired)
		return
	}

	newTokens, err := RefreshTokens(refreshToken)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, newTokens)
}

func HashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)
	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}
	return hashSalt, nil
}
