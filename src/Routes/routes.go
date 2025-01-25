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
		userRoutes.GET("/getById/:id", user.GetUserByID)
		userRoutes.GET("/getByEmail/:id", user.GetUserByEmail)
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.PUT("/update/:id", user.UpdateUser)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		adminRoutes.GET("/user/list", user.ListAllUsers)
	}
}
