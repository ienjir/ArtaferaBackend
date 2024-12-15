package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ienjir/ArtaferaBackend/controller"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/", controllers.ProtectedHandler)
		authRoutes.POST("/", controllers.LoginHandler)
	}
}
