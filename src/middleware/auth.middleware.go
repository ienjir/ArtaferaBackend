package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strings"
)

func RoleAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Authorization header is required",
			})
			return
		}

		// Check if the header starts with "Bearer "
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid authorization header format",
			})
			return
		}

		// Verify the access token
		token, serviceErr := auth.VerifyAccessToken(bearerToken[1])
		if serviceErr != nil {
			c.AbortWithStatusJSON(serviceErr.StatusCode, serviceErr)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt2.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token claims",
			})
			return
		}

		// Get user role from claims
		userRole, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Role not found in token",
			})
			return
		}

		// Check if user role is in allowed roles
		roleAllowed := false
		for _, role := range allowedRoles {
			if role == userRole {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "User role not authorized for this route",
			})
			return
		}

		// In your middleware:
		fmt.Printf("ID type: %T, value: %v\n", claims["id"], claims["id"])

		// Store user information in context for later use
		c.Set("userID", claims["id"])
		c.Set("userEmail", claims["email"])
		c.Set("userRole", userRole)
		c.Next()
	}
}
