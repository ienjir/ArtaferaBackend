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

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.RoleAuthMiddleware("admin", "user"))
	{
		userRoutes.GET("/:id", user.GetUserByID)
		userRoutes.POST("/getByEmail", user.GetUserByEmail)
		userRoutes.POST("/", user.CreateUser)
		userRoutes.PUT("/:id", user.UpdateUser)
		userRoutes.DELETE("/:id", user.DeleteUser)

		userAdminRoutes := userRoutes.Group("")
		userAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			userAdminRoutes.POST("/list", user.ListAllUsers)
		}
	}

	roleRoutes := router.Group("/role")
	roleRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		roleRoutes.GET("/:id", role.GetRoleByID)
		roleRoutes.POST("/list", role.ListRoles)
		roleRoutes.POST("/", role.CreateRole)
		roleRoutes.PUT("/:id", role.UpdateRole)
		roleRoutes.DELETE("/:id", role.DeleteRole)
	}

	languageRoutes := router.Group("/language")
	languageRoutes.Use(middleware.RoleAuthMiddleware("admin"))
	{
		languageRoutes.GET("/:id", language.GetLanguageByID)
		languageRoutes.POST("/list", language.ListLanguages)
		languageRoutes.POST("/", language.CreateLanguage)
		languageRoutes.PUT("/:id", language.UpdateLanguage)
		languageRoutes.DELETE("/:id", language.DeleteLanguage)
	}

	orderRoutes := router.Group("/order")
	orderRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		orderRoutes.GET("/:id", order.GetOrderByID)
		orderRoutes.GET("/user/:id", order.GetOrdersForUser)
		orderRoutes.POST("/list", order.ListOrder)
		orderRoutes.POST("/", order.CreateOrder)
		orderRoutes.PUT("/:id", order.UpdateOrder)

		orderAdminRoutes := orderRoutes.Group("")
		orderAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			orderAdminRoutes.DELETE("/admin/:id", order.DeleteOrder)
		}
	}

	savedRoutes := router.Group("/saved")
	savedRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		savedRoutes.GET("/:id", saved.GetSavedByID)
		savedRoutes.GET("/user/:id", saved.GetSavedForUser)
		savedRoutes.POST("/", saved.CreateSaved)
		savedRoutes.PUT("/:id", saved.UpdateSaved)
		savedRoutes.DELETE("/:id", saved.DeleteSaved)

		savedAdminRoutes := savedRoutes.Group("")
		savedAdminRoutes.Use(middleware.RoleAuthMiddleware("admin"))
		{
			savedAdminRoutes.POST("/list", saved.ListOrder)
		}
	}

	pictureRoutes := router.Group("/picture")
	pictureRoutes.Use(middleware.RoleAuthMiddleware("user", "admin"))
	{
		pictureRoutes.GET("/:id", picture.GetPictureByID)
		pictureRoutes.POST("/name", picture.GetPictureByName)
		pictureRoutes.POST("/list", picture.ListPicture)
		pictureRoutes.POST("/", picture.CreatePicture)
		pictureRoutes.PUT("/:id", picture.UpdatePicture)
		pictureRoutes.DELETE("/:id")
	}
}
