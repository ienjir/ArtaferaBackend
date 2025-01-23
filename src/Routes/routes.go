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
		authRoutes.GET("/user/getByEmail", user.GetUserByID)
	}

	// User routes
	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.RoleAuthMiddleware("user"))
	{
		// Public routes
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.GET("/test", test)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)
	}

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		adminRoutes.DELETE("/user/delete/:id", user.DeleteUser)
		adminRoutes.GET("/user/list", user.ListAllUsers)
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
