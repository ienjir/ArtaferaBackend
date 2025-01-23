package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"github.com/ienjir/ArtaferaBackend/src/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/refresh", auth.RefreshTokenHandler)
	}

	// User routes
	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)
		userRoutes.GET("/getById/:id", user.GetUserByID)
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		adminRoutes.DELETE("/user/delete/:id", user.DeleteUser)
		adminRoutes.GET("/user/list", user.ListAllUsers)
		adminRoutes.GET("/user/getById/:id", user.GetUserByID)
	}
}
