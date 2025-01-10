package roles

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
)

type RolesService struct {
	db *gorm.DB
}

func NewRolesService(db *gorm.DB) *RolesService {
	return &RolesService{db: db}
}

func (s *RolesService) Create(role *models.Role) error {
	// Hash password before saving
	result := s.db.Create(role)
	return result.Error
}
