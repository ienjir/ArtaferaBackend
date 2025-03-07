package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/api/language"
	"github.com/ienjir/ArtaferaBackend/src/api/order"
	picture "github.com/ienjir/ArtaferaBackend/src/api/picture"
	"github.com/ienjir/ArtaferaBackend/src/api/role"
	"github.com/ienjir/ArtaferaBackend/src/api/saved"
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
		languageRoutes.PUT("/update/:id", language.UpdateLanguage)
		languageRoutes.DELETE("/delete/:id", language.DeleteLanguage)
	}

	orderRoutes := router.Group("/order")
	orderRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		orderRoutes.GET("/getByID/:id", order.GetOrderByID)
		orderRoutes.GET("/getForUser/:id", order.GetOrdersForUser)
		orderRoutes.GET("/list", order.ListOrder)
		orderRoutes.POST("/create", order.CreateOrder)
		orderRoutes.PUT("/update/:id", order.UpdateOrder)

		orderAdminRoutes := orderRoutes.Group("")
		orderAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			orderAdminRoutes.DELETE("/delete/:id", order.DeleteOrder)
		}
	}

	savedRoutes := router.Group("/saved")
	savedRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		savedRoutes.GET("/getByID/:id", saved.GetSavedByID)
		savedRoutes.GET("/getForUser/:id", saved.GetSavedForUser)
		savedRoutes.POST("/create", saved.CreateSaved)
		savedRoutes.PUT("/update/:id", saved.UpdateSaved)
		savedRoutes.DELETE("/delete/:id", saved.DeleteSaved)

		savedAdminRoutes := savedRoutes.Group("")
		savedAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			savedAdminRoutes.GET("/list", saved.ListOrder)
		}
	}

	pictureRoutes := router.Group("/picture")
	pictureRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		pictureRoutes.GET("/get/:id", picture.GetPictureByID)
		pictureRoutes.POST("/create", picture.CreatePicture)
	}
}
