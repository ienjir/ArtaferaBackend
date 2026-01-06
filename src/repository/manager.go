package repository

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
)

type RepositoryManager struct {
	User           Repository[models.User]
	Role           Repository[models.Role]
	Art            ArtRepository
	ArtTranslation Repository[models.ArtTranslation]
	ArtPicture     Repository[models.ArtPicture]
	Picture        Repository[models.Picture]
	Language       Repository[models.Language]
	Currency       Repository[models.Currency]
	Order          Repository[models.Order]
	Saved          Repository[models.Saved]
	Translation    Repository[models.Translation]
	ContactMessage Repository[models.ContactMessage]
}

func NewRepositoryManager(db *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		User:           NewRepository[models.User](db),
		Role:           NewRepository[models.Role](db),
		Art:            NewArtRepository(db),
		ArtTranslation: NewRepository[models.ArtTranslation](db),
		ArtPicture:     NewRepository[models.ArtPicture](db),
		Picture:        NewRepository[models.Picture](db),
		Language:       NewRepository[models.Language](db),
		Currency:       NewRepository[models.Currency](db),
		Order:          NewRepository[models.Order](db),
		Saved:          NewRepository[models.Saved](db),
		Translation:    NewRepository[models.Translation](db),
		ContactMessage: NewRepository[models.ContactMessage](db),
	}
}
