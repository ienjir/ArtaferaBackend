package role

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func GetRoleByID(c *gin.Context) {
	var json models.GetRoleByIDRequest

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "RoleID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyGetRoleByID(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	role, err := getRoleByIDService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
	return
}

func ListRoles(c *gin.Context) {
	var json models.ListRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListRoles(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	roles, count, err := listRolesService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "roles": roles})
	return
}

func CreateRole(c *gin.Context) {
	var json models.CreateRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateRole(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.StatusCode})
		return
	}

	createdRole, err := createRoleService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": createdRole})
	return
}

func UpdateRole(c *gin.Context) {
	var json models.UpdateRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "RoleID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyUpdateRole(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	updatedRole, err := updateRoleService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": updatedRole})
	return
}

func DeleteRole(c *gin.Context) {
	var json models.DeleteRoleRequest

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "RoleID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyDeleteRole(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if err := deleteRoleService(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role successfully deleted"})
	return
}
