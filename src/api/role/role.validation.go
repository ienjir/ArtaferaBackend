package role

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetRoleByID(data models.GetRoleByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.RoleID, "RoleID").
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyListRoles(data models.ListRoleRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateOffset(int64(data.Offset)).
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyCreateRole(data models.CreateRoleRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateNotEmpty(&data.Role, "Role").
		GetFirstError()
}

func verifyUpdateRole(data models.UpdateRoleRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateNotEmpty(&data.Role, "Role").
		ValidateID(data.RoleID, "RoleID").
		GetFirstError()
}

func verifyDeleteRole(data models.DeleteRoleRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.RoleID, "RoleID").
		GetFirstError()
}
