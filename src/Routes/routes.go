package Routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/")
	}
}
