package roles

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"io"
	"net/http"
)

type RolesController struct {
	service *RolesService
}

func NewRolesController(service *RolesService) *RolesController {
	return &RolesController{service: service}
}

func (c *RolesController) Create(ctx *gin.Context) {
	var req models.Role
	rawBody, _ := io.ReadAll(ctx.Request.Body)

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := &models.Role{
		Role: req.Role,
	}

	if err := c.service.Create(role); err != nil {
		fmt.Println("Error saving role: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}
