package role

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getRoleByIDService(targetRoleID string) (*models.Role, *models.ServiceError) {
	var role models.Role

	if err := database.DB.First(&role, targetRoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Role not found"}
		} else {
			return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving role"}
		}
	}

	return &role, nil
}

func listRolesService(offset int) (*[]models.Role, *int64, *models.ServiceError) {
	var roles []models.Role
	var count int64

	if err := database.DB.Limit(5).Offset(offset * 10).Find(&roles).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving users from database",
		}
	}

	if err := database.DB.Model(&models.Role{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting users in database",
		}
	}

	return &roles, &count, nil
}

func createRoleService(request models.CreateRoleRequest) (*models.Role, *models.ServiceError) {
	var role models.Role
	var newRole models.Role

	// Check if role already exists
	if err := database.DB.Where("role = ?", request.Role).First(&role).Error; err == nil {
		return nil, &models.ServiceError{StatusCode: http.StatusConflict, Message: "Role already exists"}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Database error"}
	}

	newRole = models.Role{
		Role: request.Role,
	}

	if err := database.DB.Create(&newRole).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to save user"}
	}

	return &newRole, nil
}
