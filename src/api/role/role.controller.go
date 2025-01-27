package role

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func GetRoleByID(c *gin.Context) {
	requestUserID := c.GetInt64("userID")
	requestUserRole := c.GetString("userRole")
	targetRoleID := c.Param("id")

	if err := verifyGetRoleByID(requestUserID, requestUserRole, targetRoleID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	role, err := getRoleByIDService(targetRoleID)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
	return
}

func ListRoles(c *gin.Context) {
	var Data models.ListRoleRequest

	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles, count, err := listRolesService(Data.Offset)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "roles": roles})
	return
}

func CreateRole(c *gin.Context) {
	var role models.CreateRoleRequest

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRole, err := createRoleService(role)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": createdRole})
	return
}

func UpdateRole(c *gin.Context) {
	var role models.UpdateRoleRequest

	role.RoleID = c.Param("id")

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRole, err := updateRoleService(role)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": updatedRole})
	return
}
