package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, Database *gorm.DB) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/", auth.ProtectedHandler)
		authRoutes.POST("/", auth.LoginHandler)
	}

	// User Routes
	userService := user.NewUserService(Database)
	userController := user.NewUserController(userService)
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
	}
}
