package middleware

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/utils"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
				Error: utils.ErrAccessTokenRequired,
				Code:  http.StatusUnauthorized,
			})
			return
		}

		// Check if the header starts with "Bearer"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
				Error: utils.ErrInvalidAuthHeader,
				Code:  http.StatusUnauthorized,
			})
			return
		}

		token, serviceErr := auth.VerifyAccessToken(bearerToken[1])
		if serviceErr != nil {
			c.AbortWithStatusJSON(serviceErr.StatusCode, utils.ErrorResponse{
				Error: serviceErr.Message,
				Code:  serviceErr.StatusCode,
			})
			return
		}

		claims, ok := token.Claims.(jwt2.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
				Error: utils.ErrInvalidTokenClaims,
				Code:  http.StatusUnauthorized,
			})
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
				Error: utils.ErrRoleNotInToken,
				Code:  http.StatusUnauthorized,
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
			c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorResponse{
				Error: utils.ErrInsufficientPerms,
				Code:  http.StatusForbidden,
			})
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
