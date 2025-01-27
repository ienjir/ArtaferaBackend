package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/language"
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
	userRoutes.Use(middleware.RoleAuthMiddleware("admin", "user"))
	{
		userRoutes.GET("/getByID/:id", user.GetUserByID)
		userRoutes.GET("/getByEmail", user.GetUserByEmail)
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.PUT("/update/:id", user.UpdateUser)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)

		userAdminRoutes := userRoutes.Group("")
		userAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			userAdminRoutes.GET("/list", user.ListAllUsers)
		}
	}

	roleRoutes := router.Group("/role")
	roleRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		roleRoutes.GET("/getByID/:id", role.GetRoleByID)
		roleRoutes.GET("/list", role.ListRoles)
		roleRoutes.POST("/create", role.CreateRole)
		roleRoutes.PUT("/update/:id", role.UpdateRole)
		roleRoutes.DELETE("/delete/:id", role.DeleteRole)
	}

	languageRoutes := router.Group("/language")
	languageRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		languageRoutes.GET("/getByID/:id", language.GetLanguageByID)
		languageRoutes.GET("/list", language.ListLanguages)
		languageRoutes.POST("/create", language.CreateLanguage)
	}
}
