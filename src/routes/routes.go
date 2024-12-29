package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
)

func RegisterRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/", auth.ProtectedHandler)
		authRoutes.POST("/", auth.LoginHandler)
	}

}
