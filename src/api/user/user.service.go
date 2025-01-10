// user.service.go
package user

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(user *models.User) error {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result := s.db.Create(user)
	return result.Error
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) Update(user *models.User) error {
	// Don't update password if it's empty
	if len(user.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	} else {
		// Exclude password from update
		return s.db.Model(user).Omit("password").Updates(user).Error
	}

	return s.db.Save(user).Error
}

func (s *UserService) Delete(id uint) error {
	// Soft delete
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
}

func (s *UserService) List(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Get total count
	if err := s.db.Model(&models.User{}).Where("is_deleted = ?", false).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	result := s.db.Where("is_deleted = ?", false).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}
