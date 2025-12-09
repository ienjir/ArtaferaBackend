package repository

import (
	"errors"

	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

type ArtRepository interface {
	Repository[models.Art]
	GetPublicArtByID(id int64, languageCode string) (*models.Art, *models.ServiceError)
	ListPublicArts(languageCode string, offset, limit int) (*[]models.Art, *models.ServiceError)
	CountPublicArts() (*int64, *models.ServiceError)
}

type GormArtRepository struct {
	*GormRepository[models.Art]
}

func NewArtRepository(db *gorm.DB) ArtRepository {
	return &GormArtRepository{
		GormRepository: &GormRepository[models.Art]{db: db},
	}
}

func (r *GormArtRepository) GetPublicArtByID(id int64, languageCode string) (*models.Art, *models.ServiceError) {
	var art models.Art
	query := r.db.Where("visible = ?", true)

	// Preload pictures with their picture data, ordered by priority
	query = query.Preload("Pictures", func(db *gorm.DB) *gorm.DB {
		return db.Order("priority ASC")
	})
	query = query.Preload("Pictures.Picture")

	// Preload currency
	query = query.Preload("Currency")

	// If language code provided, filter translations; otherwise load all
	if languageCode != "" {
		query = query.Preload("Translations", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN languages ON languages.id = art_translations.language_id").
				Where("languages.language_code = ?", languageCode)
		})
	} else {
		query = query.Preload("Translations")
	}
	query = query.Preload("Translations.Language")

	if err := query.First(&art, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	// Post-process: filter out non-public pictures
	var publicPictures []models.ArtPicture
	for _, pic := range art.Pictures {
		if pic.Picture.IsPublic {
			publicPictures = append(publicPictures, pic)
		}
	}
	art.Pictures = publicPictures

	return &art, nil
}

func (r *GormArtRepository) ListPublicArts(languageCode string, offset, limit int) (*[]models.Art, *models.ServiceError) {
	var arts []models.Art

	// Get all visible arts (including sold ones for portfolio)
	query := r.db.Where("visible = ?", true).
		Offset(offset * 20).
		Limit(20)

	// Preload currency (simple 1:1, no issue)
	query = query.Preload("Currency")

	// Preload ALL pictures first, then we'll filter in Go
	query = query.Preload("Pictures", func(db *gorm.DB) *gorm.DB {
		return db.Order("priority ASC")
	})
	query = query.Preload("Pictures.Picture")

	// Preload ALL translations, then we'll filter in Go
	if languageCode != "" {
		query = query.Preload("Translations", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN languages ON languages.id = art_translations.language_id").
				Where("languages.language_code = ?", languageCode)
		})
	} else {
		query = query.Preload("Translations")
	}
	query = query.Preload("Translations.Language")

	if err := query.Find(&arts).Error; err != nil {
		return nil, utils.NewDatabaseRetrievalError()
	}

	// Post-process: keep only the first PUBLIC picture and first translation per art
	for i := range arts {
		// Filter out non-public pictures and keep only the first public one
		var publicPictures []models.ArtPicture
		for _, pic := range arts[i].Pictures {
			if pic.Picture.IsPublic {
				publicPictures = append(publicPictures, pic)
				break // Only take the first public picture
			}
		}
		arts[i].Pictures = publicPictures

		// Keep only first translation
		if len(arts[i].Translations) > 1 {
			arts[i].Translations = arts[i].Translations[:1]
		}
	}

	return &arts, nil
}

func (r *GormArtRepository) CountPublicArts() (*int64, *models.ServiceError) {
	var count int64
	if err := r.db.Model(&models.Art{}).Where("visible = ?", true).Count(&count).Error; err != nil {
		return nil, utils.NewDatabaseCountError()
	}
	return &count, nil
}
