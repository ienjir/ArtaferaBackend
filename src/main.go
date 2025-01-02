// routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/", auth.ProtectedHandler)
		authRoutes.POST("/", auth.LoginHandler)
	}

	// User routes
	userService := user.NewUserService(db)
	userController := user.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.Create)
		userRoutes.GET("/:id", userController.Get)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
		userRoutes.GET("/", userController.List)
	}
}
