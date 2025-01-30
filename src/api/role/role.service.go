package role

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getRoleByIDService(data models.GetRoleByIDRequest) (*models.Role, *models.ServiceError) {
	var role models.Role

	if err := database.DB.First(&role, data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Role not found"}
		} else {
			return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving role"}
		}
	}

	return &role, nil
}

func listRolesService(data models.ListRoleRequest) (*[]models.Role, *int64, *models.ServiceError) {
	var roles []models.Role
	var count int64

	if err := database.DB.Limit(5).Offset(int(data.Offset) * 10).Find(&roles).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving roles from database",
		}
	}

	if err := database.DB.Model(&models.Role{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting roles in database",
		}
	}

	return &roles, &count, nil
}

func createRoleService(data models.CreateRoleRequest) (*models.Role, *models.ServiceError) {
	var role models.Role
	var newRole models.Role

	if err := database.DB.Where("name = ?", data.Role).First(&role).Error; err == nil {
		return nil, &models.ServiceError{StatusCode: http.StatusConflict, Message: "Role already exists"}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Database error",
		}
	}

	newRole = models.Role{
		Name: data.Role,
	}

	if err := database.DB.Create(&newRole).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save role",
		}
	}

	return &newRole, nil
}

func updateRoleService(data models.UpdateRoleRequest) (*models.Role, *models.ServiceError) {
	var role models.Role

	if err := database.DB.First(&role, "id = ?", data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Role not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	role.Name = data.Role

	if err := database.DB.Save(&role).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Error while updating role",
		}
	}

	return &role, nil
}

func deleteRoleService(data models.DeleteRoleRequest) *models.ServiceError {
	var role models.Role

	if err := database.DB.First(&role, "id = ?", data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Role not found",
			}
		}
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if result := database.DB.Delete(&models.Role{}, data.RoleID); result.Error != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while deleting role",
		}
	}

	return nil
}
