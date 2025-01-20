package middleware

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strings"
)

// RoleAuthMiddleware creates a middleware that checks if the user's role is allowed
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

		// Verify the token
		tokenString := bearerToken[1]
		if err := auth.VerifyToken(tokenString); err != nil {
			c.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		// Parse the token to get claims
		token, _ := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
			return auth.JWTSecret, nil
		})

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

		// Store user information in context for later use
		c.Set("userID", claims["id"])
		c.Set("userEmail", claims["email"])
		c.Set("userRole", userRole)

		c.Next()
	}
}
