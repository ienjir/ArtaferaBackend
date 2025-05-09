package middleware

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strings"
)

func RoleAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, role := range allowedRoles {
			if role == "all" {
				c.Next()
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Access token is required",
			})
			return
		}

		// Check if the header starts with "Bearer"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid authorization header format",
			})
			return
		}

		token, serviceErr := auth.VerifyAccessToken(bearerToken[1])
		if serviceErr != nil {
			c.AbortWithStatusJSON(serviceErr.StatusCode, gin.H{"error": serviceErr.Message})
			return
		}

		claims, ok := token.Claims.(jwt2.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token claims",
			})
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Role not found in token",
			})
			return
		}

		userID, ok := claims["id"].(float64)

		roleAllowed := false
		for _, role := range allowedRoles {
			if role == userRole {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User is not authorized for this route"})
			return
		}

		userIDInt := int64(userID)

		// Store user info in context
		c.Set("userID", userIDInt)
		c.Set("userEmail", claims["email"])
		c.Set("userRole", userRole)
		c.Next()
	}
}
