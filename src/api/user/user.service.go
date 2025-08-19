package user

import (
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"strings"
)

func getUserByIDService(data models.GetUserByIDRequest) (*models.User, *models.ServiceError) {
	user, err := database.Repositories.User.GetByID(data.TargetID, "Role")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, err
	}

	return user, nil
}

func getUserByEmailService(data models.GetUserByEmailRequest) (*models.User, *models.ServiceError) {
	data.Email = strings.ToLower(data.Email)

	user, err := database.Repositories.User.FindByField("email", data.Email, "Role")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, err
	}

	if data.UserRole != "admin" {
		if int(data.UserID) != int(user.ID) {
			return nil, utils.NewOwnerAccessError()
		}
	}

	return user, nil
}

func listUsersService(data models.ListUserRequest) (*[]models.User, *int64, *models.ServiceError) {
	users, err := database.Repositories.User.List(int(data.Offset*10), 10, "Role")
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.User.Count()
	if err != nil {
		return nil, nil, err
	}

	return users, count, nil
}

func createUserService(data models.CreateUserRequest) (*models.User, *models.ServiceError) {
	data.Email = strings.ToLower(data.Email)

	// Check if user already exists
	if existingUser, err := database.Repositories.User.FindByField("email", data.Email); err == nil && existingUser != nil {
		return nil, utils.NewUserAlreadyExistsError()
	}

	hashedPassword, err := auth.HashPassword(data.Password)
	if err != nil {
		return nil, utils.NewHashPasswordError()
	}

	newUser := models.User{
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Email:       data.Email,
		Phone:       data.Phone,
		PhoneRegion: data.PhoneRegion,
		Address1:    data.Address1,
		Address2:    data.Address2,
		City:        data.City,
		PostalCode:  data.PostalCode,
		Password:    hashedPassword.Hash,
		Salt:        hashedPassword.Salt,
	}

	if serviceErr := database.Repositories.User.Create(&newUser); serviceErr != nil {
		return nil, serviceErr
	}

	return &newUser, nil
}

func updateUserService(data models.UpdateUserRequest) (*models.User, *models.ServiceError) {
	user, err := database.Repositories.User.GetByID(data.TargetID, "Role")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, err
	}

	if data.Firstname != nil {
		user.Firstname = *data.Firstname
	}
	if data.Lastname != nil {
		user.Lastname = *data.Lastname
	}
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.Phone != nil {
		user.Phone = data.Phone
	}
	if data.PhoneRegion != nil {
		user.PhoneRegion = data.PhoneRegion
	}
	if data.Address1 != nil {
		user.Address1 = data.Address1
	}
	if data.Address2 != nil {
		user.Address2 = data.Address2
	}
	if data.City != nil {
		user.City = data.City
	}
	if data.PostalCode != nil {
		user.PostalCode = data.PostalCode
	}

	if data.Password != nil {
		password, hashErr := auth.HashPassword(*data.Password)
		if hashErr != nil {
			return nil, utils.NewHashPasswordError()
		}
		user.Password = password.Hash
		user.Salt = password.Salt
	}

	if serviceErr := database.Repositories.User.Update(user); serviceErr != nil {
		return nil, serviceErr
	}

	return user, nil
}

func deleteUserService(data models.DeleteUserRequest) *models.ServiceError {
	if err := database.Repositories.User.Delete(data.TargetID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewUserNotFoundError()
		}
		return err
	}

	return nil
}
