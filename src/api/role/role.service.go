package role

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

func getRoleByIDService(data models.GetRoleByIDRequest) (*models.Role, *models.ServiceError) {
	var role models.Role

	if err := database.DB.First(&role, data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRoleNotFoundError()
		} else {
			return nil, utils.NewDatabaseRetrievalError()
		}
	}

	return &role, nil
}

func listRolesService(data models.ListRoleRequest) (*[]models.Role, *int64, *models.ServiceError) {
	var roles []models.Role
	var count int64

	if err := database.DB.Limit(5).Offset(int(data.Offset) * 10).Find(&roles).Error; err != nil {
		return nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Role{}).Count(&count).Error; err != nil {
		return nil, nil, utils.NewDatabaseCountError()
	}

	return &roles, &count, nil
}

func createRoleService(data models.CreateRoleRequest) (*models.Role, *models.ServiceError) {
	var role models.Role
	var newRole models.Role

	if err := database.DB.Where("name = ?", data.Role).First(&role).Error; err == nil {
		return nil, utils.NewRoleAlreadyExistsError()
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewDatabaseRetrievalError()
	}

	newRole = models.Role{
		Name: data.Role,
	}

	if err := database.DB.Create(&newRole).Error; err != nil {
		return nil, utils.NewDatabaseCreateError()
	}

	return &newRole, nil
}

func updateRoleService(data models.UpdateRoleRequest) (*models.Role, *models.ServiceError) {
	var role models.Role

	if err := database.DB.First(&role, "id = ?", data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRoleNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	role.Name = data.Role

	if err := database.DB.Save(&role).Error; err != nil {
		return nil, utils.NewDatabaseUpdateError()
	}

	return &role, nil
}

func deleteRoleService(data models.DeleteRoleRequest) *models.ServiceError {
	var role models.Role

	if err := database.DB.First(&role, "id = ?", data.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewRoleNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if result := database.DB.Delete(&models.Role{}, data.RoleID); result.Error != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
}
