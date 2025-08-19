package art

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getArtByIDService(data models.GetArtByIDRequest) (*models.Art, *models.ServiceError) {
	var Art models.Art

	if err := database.DB.First(&Art, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
			}
		} else {
			return nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving Art",
			}
		}
	}

	return &Art, nil
}

func listArtService(data models.ListArtRequest) (*[]models.Art, *int64, *models.ServiceError) {
	var arts []models.Art
	var count int64

	query := database.DB.Model(&models.Art{})

	if data.Available != nil {
		query = query.Where("available = ?", *data.Available)
	}

	if data.UserRole != "admin" {
		query = query.Where("visible = ?", true)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting arts",
		}
	}

	offset := (data.Page - 1) * data.PageSize
	sortBy := "created_at"
	sortOrder := "desc"

	if data.SortBy != nil {
		sortBy = *data.SortBy
	}
	if data.SortOrder != nil {
		sortOrder = *data.SortOrder
	}

	orderClause := sortBy + " " + sortOrder

	if err := query.Preload("Currency").Preload("Pictures").Preload("Translations").
		Offset(offset).Limit(data.PageSize).Order(orderClause).Find(&arts).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving arts",
		}
	}

	return &arts, &count, nil
}

func createArtService(data models.CreateArtRequest) (*models.Art, *models.ServiceError) {
	art := models.Art{
		Price:        data.Price,
		CurrencyID:   data.CurrencyID,
		CreationYear: data.CreationYear,
		Width:        data.Width,
		Height:       data.Height,
		Depth:        data.Depth,
	}

	if data.Available != nil {
		art.Available = *data.Available
	} else {
		art.Available = true
	}

	if data.Visible != nil {
		art.Visible = *data.Visible
	} else {
		art.Visible = true
	}

	if err := database.DB.Create(&art).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while creating art",
		}
	}

	return &art, nil
}

func updateArtService(data models.UpdateArtRequest) (*models.Art, *models.ServiceError) {
	var art models.Art

	if err := database.DB.First(&art, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving art",
		}
	}

	updates := make(map[string]interface{})

	if data.Price != nil {
		updates["price"] = *data.Price
	}
	if data.CurrencyID != nil {
		updates["currency_id"] = *data.CurrencyID
	}
	if data.CreationYear != nil {
		updates["creation_year"] = *data.CreationYear
	}
	if data.Width != nil {
		updates["width"] = *data.Width
	}
	if data.Height != nil {
		updates["height"] = *data.Height
	}
	if data.Depth != nil {
		updates["depth"] = *data.Depth
	}
	if data.Available != nil {
		updates["available"] = *data.Available
	}
	if data.Visible != nil {
		updates["visible"] = *data.Visible
	}

	if err := database.DB.Model(&art).Updates(updates).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while updating art",
		}
	}

	return &art, nil
}

func deleteArtService(data models.DeleteArtRequest) *models.ServiceError {
	var art models.Art

	if err := database.DB.First(&art, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
			}
		}
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving art",
		}
	}

	if err := database.DB.Delete(&art).Error; err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while deleting art",
		}
	}

	return nil
}
