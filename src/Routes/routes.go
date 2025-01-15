package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/", user.CreateUser)
	}
}
