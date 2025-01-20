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
	}

	// User routes
	userRoutes := router.Group("/user")
	{
		// Public routes
		userRoutes.POST("/create", user.CreateUser)

		// Protected user routes (example)
		protected := userRoutes.Group("")
		protected.Use(middleware.RoleAuthMiddleware("admin", "user"))
		{
			// Add your protected user routes here
			// protected.GET("/profile", user.GetProfile)
			// protected.PUT("/update", user.UpdateUser)
		}
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		// Add your admin routes here
		// adminRoutes.GET("/users", admin.GetAllUsers)
		// adminRoutes.DELETE("/users/:id", admin.DeleteUser)
	}
}
