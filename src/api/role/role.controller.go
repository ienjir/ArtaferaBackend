package role

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func GetRoleByID(c *gin.Context) {
	var json models.GetRoleByIDRequest

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyGetRoleByID(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	role, err := getRoleByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"role": role})
	return
}

func ListRoles(c *gin.Context) {
	var json models.ListRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListRoles(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	roles, count, err := listRolesService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "roles": roles})
	return
}

func CreateRole(c *gin.Context) {
	var json models.CreateRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateRole(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	createdRole, err := createRoleService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"role": createdRole})
	return
}

func UpdateRole(c *gin.Context) {
	var json models.UpdateRoleRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyUpdateRole(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	updatedRole, err := updateRoleService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"role": updatedRole})
	return
}

func DeleteRole(c *gin.Context) {
	var json models.DeleteRoleRequest

	roleID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.RoleID = roleID

	if err := verifyDeleteRole(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteRoleService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil, "Role successfully deleted")
	return
}
