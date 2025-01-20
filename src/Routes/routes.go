package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/user"
	"github.com/ienjir/ArtaferaBackend/src/middleware"
	"net/http"
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
		userRoutes.GET("/test", test)

		// Protected user routes (example)
		protected := userRoutes.Group("")
		protected.Use(middleware.RoleAuthMiddleware("admin"))
		{
			// Add your protected user routes here
			protected.GET("/test2", test2)
			// protected.GET("/profile", user.GetProfile)
			// protected.PUT("/update", user.UpdateUser)
		}
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		adminRoutes.GET("/test", test)
		// Add your admin routes here
		// adminRoutes.GET("/users", admin.GetAllUsers)
		// adminRoutes.DELETE("/users/:id", admin.DeleteUser)
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
	return
}

func test2(c *gin.Context) {
	c.JSON(http.StatusOK, "test2")
	return
}
