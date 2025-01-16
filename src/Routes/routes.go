package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/create", user.CreateUser)
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", auth.LoginHandler)
	}
}
