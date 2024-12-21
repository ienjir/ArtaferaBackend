package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/files"
)

func RegisterRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/", auth.ProtectedHandler)
		authRoutes.POST("/", auth.LoginHandler)
	}

	filesRoutes := router.Group("/files")
	{
		filesRoutes.POST("/", files.UploadFile)
		filesRoutes.POST("/d§", files.UploadSingleFile)
	}
}
