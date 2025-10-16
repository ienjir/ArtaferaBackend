package repository

import (
	"errors"
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) *models.ServiceError
	GetByID(id int64, preloads ...string) (*T, *models.ServiceError)
	Update(entity *T) *models.ServiceError
	UpdateFields(id int64, updates map[string]interface{}) (*T, *models.ServiceError)
	Delete(id int64) *models.ServiceError
	DeleteEntity(entity *T) *models.ServiceError
	List(offset, limit int, preloads ...string) (*[]T, *models.ServiceError)
	Count() (*int64, *models.ServiceError)
	FindByField(field string, value interface{}, preloads ...string) (*T, *models.ServiceError)
	FindAllByField(field string, value interface{}, offset, limit int, preloads ...string) (*[]T, *models.ServiceError)
	Query() *gorm.DB
}

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) Create(entity *T) *models.ServiceError {
	if err := r.db.Create(entity).Error; err != nil {
		return utils.NewDatabaseCreateError()
	}
	return nil
}

func (r *GormRepository[T]) GetByID(id int64, preloads ...string) (*T, *models.ServiceError) {
	var entity T
	query := r.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRecordNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	return &entity, nil
}

func (r *GormRepository[T]) Update(entity *T) *models.ServiceError {
	if err := r.db.Save(entity).Error; err != nil {
		return utils.NewDatabaseUpdateError()
	}
	return nil
}

func (r *GormRepository[T]) UpdateFields(id int64, updates map[string]interface{}) (*T, *models.ServiceError) {
	var entity T
	
	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRecordNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if err := r.db.Model(&entity).Updates(updates).Error; err != nil {
		return nil, utils.NewDatabaseUpdateError()
	}

	return &entity, nil
}

func (r *GormRepository[T]) Delete(id int64) *models.ServiceError {
	var entity T
	
	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewRecordNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if err := r.db.Delete(&entity, id).Error; err != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
}

func (r *GormRepository[T]) DeleteEntity(entity *T) *models.ServiceError {
	if err := r.db.Delete(entity).Error; err != nil {
		return utils.NewDatabaseDeleteError()
	}
	return nil
}

func (r *GormRepository[T]) List(offset, limit int, preloads ...string) (*[]T, *models.ServiceError) {
	var entities []T
	query := r.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		return nil, utils.NewDatabaseRetrievalError()
	}

	return &entities, nil
}

func (r *GormRepository[T]) Count() (*int64, *models.ServiceError) {
	var count int64
	var entity T

	if err := r.db.Model(&entity).Count(&count).Error; err != nil {
		return nil, utils.NewDatabaseCountError()
	}

	return &count, nil
}

func (r *GormRepository[T]) FindByField(field string, value interface{}, preloads ...string) (*T, *models.ServiceError) {
	var entity T
	query := r.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where(fmt.Sprintf("%s = ?", field), value).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRecordNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	return &entity, nil
}

func (r *GormRepository[T]) FindAllByField(field string, value interface{}, offset, limit int, preloads ...string) (*[]T, *models.ServiceError) {
	var entities []T
	query := r.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where(fmt.Sprintf("%s = ?", field), value).Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		return nil, utils.NewDatabaseRetrievalError()
	}

	return &entities, nil
}

func (r *GormRepository[T]) Query() *gorm.DB {
	return r.db
}