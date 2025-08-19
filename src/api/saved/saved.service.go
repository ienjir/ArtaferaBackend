package saved

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func getSavedByIDService(data models.GetSavedByIDRequest) (*models.Saved, *models.ServiceError) {
	saved, err := database.Repositories.Saved.GetByID(data.TargetID, "Art", "User")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewSavedNotFoundError()
		}
		return nil, err
	}

	return saved, nil
}

func getSavedForUserService(data models.GetSavedForUserRequest) (*[]models.Saved, *models.User, *int64, *models.ServiceError) {
	user, err := database.Repositories.User.GetByID(data.TargetUserID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, nil, nil, utils.NewUserNotFoundError()
		}
		return nil, nil, nil, err
	}

	saved, err := database.Repositories.Saved.FindAllByField("user_id", data.TargetUserID, int(data.Offset)*5, 5, "Art")
	if err != nil {
		return nil, nil, nil, err
	}

	// Count saved items for user
	query := database.Repositories.Saved.Query().Where("user_id = ?", data.TargetUserID)
	var count int64
	if countErr := query.Count(&count).Error; countErr != nil {
		return nil, nil, nil, utils.NewDatabaseCountError()
	}

	return saved, user, &count, nil
}

func listSavedService(data models.ListSavedRequest) (*[]models.Saved, *int64, *models.ServiceError) {
	saved, err := database.Repositories.Saved.List(int(data.Offset*5), 5, "Art", "User")
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.Saved.Count()
	if err != nil {
		return nil, nil, err
	}

	return saved, count, nil
}

func createSavedService(data models.CreateSavedRequest) (*models.Saved, *models.ServiceError) {
	// Verify art exists
	art, err := database.Repositories.Art.GetByID(data.ArtID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, err
	}

	// Verify user exists
	user, err := database.Repositories.User.GetByID(*data.TargetUserID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, err
	}

	// Check if already saved
	query := database.Repositories.Saved.Query().Where("user_id = ? AND art_id = ?", user.ID, data.ArtID)
	var existingSaved models.Saved
	if queryErr := query.First(&existingSaved).Error; queryErr == nil {
		return nil, utils.NewArtAlreadySavedError()
	}

	newSaved := models.Saved{
		UserID: user.ID,
		ArtID:  art.ID,
	}

	if err := database.Repositories.Saved.Create(&newSaved); err != nil {
		return nil, err
	}

	return &newSaved, nil
}

func updateSavedService(data models.UpdateSavedRequest) (*models.Saved, *models.ServiceError) {
	saved, err := database.Repositories.Saved.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewSavedNotFoundError()
		}
		return nil, err
	}

	if data.TargetUserID != nil {
		saved.UserID = *data.TargetUserID
	}

	if data.ArtID != nil {
		saved.ArtID = *data.ArtID
	}

	if err := database.Repositories.Saved.Update(saved); err != nil {
		return nil, err
	}

	return saved, nil
}

func deleteSavedService(data models.DeleteSavedRequest) *models.ServiceError {
	saved, err := database.Repositories.Saved.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return utils.NewSavedNotFoundError()
		}
		return err
	}

	if data.UserRole != "admin" {
		if saved.UserID != data.UserID {
			return utils.NewNotAllowedRouteError()
		}
	}

	if err := database.Repositories.Saved.DeleteEntity(saved); err != nil {
		return err
	}

	return nil
}
