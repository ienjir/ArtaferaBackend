package art

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func getArtByIDService(data models.GetArtByIDRequest) (*models.Art, *models.ServiceError) {
	art, err := database.Repositories.Art.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, err
	}

	return art, nil
}

func listArtService(data models.ListArtRequest) (*[]models.Art, *int64, *models.ServiceError) {
	artItems, err := database.Repositories.Art.List(int(data.Offset*30), 30)
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.Art.Count()
	if err != nil {
		return nil, nil, err
	}

	return artItems, count, nil
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

	if err := database.Repositories.Art.Create(&art); err != nil {
		return nil, err
	}

	return &art, nil
}

func updateArtService(data models.UpdateArtRequest) (*models.Art, *models.ServiceError) {
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

	art, err := database.Repositories.Art.UpdateFields(data.TargetID, updates)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, err
	}

	return art, nil
}

func deleteArtService(data models.DeleteArtRequest) *models.ServiceError {
	if err := database.Repositories.Art.Delete(data.TargetID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewArtNotFoundError()
		}
		return err
	}

	return nil
}
