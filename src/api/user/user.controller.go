// user.controller.go
package user

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *UserService
}

type CreateUserRequest struct {
	Firstname  string  `json:"firstname" binding:"required"`
	Lastname   string  `json:"lastname" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Phone      *string `json:"phone"`
	Address1   *string `json:"address1"`
	Address2   *string `json:"address2"`
	City       *string `json:"city"`
	PostalCode *string `json:"postal_code"`
	Password   string  `json:"password" binding:"required,min=6"`
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Email:      req.Email,
		Phone:      req.Phone,
		Address1:   req.Address1,
		Address2:   req.Address2,
		City:       req.City,
		PostalCode: req.PostalCode,
		Password:   []byte(req.Password),
	}

	if err := c.service.Create(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Clear password before sending response
	user.Password = nil
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Clear password before sending response
	user.Password = nil
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Email:      req.Email,
		Phone:      req.Phone,
		Address1:   req.Address1,
		Address2:   req.Address2,
		City:       req.City,
		PostalCode: req.PostalCode,
	}
	user.ID = uint(id)

	if req.Password != "" {
		user.Password = []byte(req.Password)
	}

	if err := c.service.Update(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Clear password before sending response
	user.Password = nil
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (c *UserController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	users, total, err := c.service.List(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Clear passwords before sending response
	for i := range users {
		users[i].Password = nil
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}
