package role

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func getRoleByIDService(data models.GetRoleByIDRequest) (*models.Role, *models.ServiceError) {
	role, err := database.Repositories.Role.GetByID(data.RoleID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewRoleNotFoundError()
		}
		return nil, err
	}

	return role, nil
}

func listRolesService(data models.ListRoleRequest) (*[]models.Role, *int64, *models.ServiceError) {
	roles, err := database.Repositories.Role.List(int(data.Offset)*10, 5)
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.Role.Count()
	if err != nil {
		return nil, nil, err
	}

	return roles, count, nil
}

func createRoleService(data models.CreateRoleRequest) (*models.Role, *models.ServiceError) {
	// Check if role already exists
	if existingRole, err := database.Repositories.Role.FindByField("name", data.Role); err == nil && existingRole != nil {
		return nil, utils.NewRoleAlreadyExistsError()
	}

	newRole := models.Role{
		Name: data.Role,
	}

	if err := database.Repositories.Role.Create(&newRole); err != nil {
		return nil, err
	}

	return &newRole, nil
}

func updateRoleService(data models.UpdateRoleRequest) (*models.Role, *models.ServiceError) {
	role, err := database.Repositories.Role.GetByID(data.RoleID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewRoleNotFoundError()
		}
		return nil, err
	}

	role.Name = data.Role

	if err := database.Repositories.Role.Update(role); err != nil {
		return nil, err
	}

	return role, nil
}

func deleteRoleService(data models.DeleteRoleRequest) *models.ServiceError {
	if err := database.Repositories.Role.Delete(data.RoleID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewRoleNotFoundError()
		}
		return err
	}

	return nil
}
