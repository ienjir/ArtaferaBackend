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
	
	query = query.Preload("Pictures", func(db *gorm.DB) *gorm.DB {
		return db.Order("COALESCE(priority, 999999)")
	})
	query = query.Preload("Pictures.Picture", "is_public = ?", true)
	
	query = query.Preload("Currency")
	
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
	
	return &art, nil
}

func (r *GormArtRepository) ListPublicArts(languageCode string, offset, limit int) (*[]models.Art, *models.ServiceError) {
	var arts []models.Art
	
	query := r.db.Where("visible = ?", true).
		Offset(offset).
		Limit(limit)
	
	query = query.Preload("Currency")
	
	query = query.Preload("Pictures", func(db *gorm.DB) *gorm.DB {
		return db.Order("COALESCE(priority, 999999)")
	})
	query = query.Preload("Pictures.Picture", "is_public = ?", true)
	
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
	
	for i := range arts {
		if len(arts[i].Pictures) > 1 {
			arts[i].Pictures = arts[i].Pictures[:1]
		}
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
