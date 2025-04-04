package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"time"
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
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := VerifyUser(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	rememberMe := json.RememberMe
	var jwt *TokenPair

	if rememberMe {
		jwt, err = GenerateTokenPair(*user, true)
	} else {
		jwt, err = GenerateTokenPair(*user, false)
	}

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	var accessTokenExpiry, refreshTokenExpiry time.Duration
	if rememberMe {
		accessTokenExpiry = time.Hour * 24 * 30
		refreshTokenExpiry = time.Hour * 24 * 30
	} else {
		accessTokenExpiry = time.Hour
		refreshTokenExpiry = time.Hour * 24
	}
	c.SetCookie("access_token", jwt.AccessToken, int(accessTokenExpiry/time.Second), "/", "", false, true)
	c.SetCookie("refresh_token", jwt.RefreshToken, int(refreshTokenExpiry/time.Second), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  jwt.AccessToken,
		"refresh_token": jwt.RefreshToken,
	})
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

func Me(c *gin.Context) {
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		tokenString = c.GetHeader("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no token provided"})
			return
		}
	}

	token, err2 := VerifyAccessToken(tokenString)
	if err2 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	var user models.User
	if err := database.DB.Preload("Role").Where("id = ?", uint(userID)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role.Name,
	})
}
