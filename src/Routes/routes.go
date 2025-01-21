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
		authRoutes.POST("/refresh", auth.RefreshTokenHandler)
	}

	// User routes
	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.RoleAuthMiddleware("user"))
	{
		// Public routes
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.GET("/test", test)

		protected := userRoutes.Group("")
		protected.Use(middleware.RoleAuthMiddleware("admin"))
		{
			protected.GET("/test2", test2)
		}
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		adminRoutes.GET("/test", test)
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
