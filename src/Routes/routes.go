package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/role"
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
	userRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		userRoutes.GET("/getByID/:id", user.GetUserByID)
		userRoutes.GET("/getByEmail/:id", user.GetUserByEmail)
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.PUT("/update/:id", user.UpdateUser)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)

		userAdminRoutes := userRoutes.Group("/")
		userAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			userAdminRoutes.Group("/list", user.ListAllUsers)
		}
	}

	roleRoutes := router.Group("/role")
	roleRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		roleRoutes.GET("/getByID/:id", role.GetRoleByID)
		roleRoutes.GET("/list", role.ListRoles)
	}
}
