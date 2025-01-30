package role

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetRoleByID(data models.GetRoleByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.RoleID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "OrderID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see this role",
		}
	}

	return nil
}

func verifyListRoles(data models.ListRoleRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to  be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see all role",
		}
	}

	return nil
}

func verifyCreateRole(data models.CreateRoleRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see all role",
		}
	}

	if data.Role == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Role can not be empty",
		}
	}

	return nil
}

func verifyUpdateRole(data models.UpdateRoleRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see all role",
		}
	}

	if data.Role == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Role can not be empty",
		}
	}

	if data.RoleID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "RoleID has to be over 1",
		}
	}

	return nil
}

func verifyDeleteRole(data models.DeleteRoleRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see all role",
		}
	}

	if data.RoleID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "RoleID has to be over 1",
		}
	}

	return nil
}
