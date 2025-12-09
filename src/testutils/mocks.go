package testutils

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
)

// MockRepository is a mock implementation of the Repository interface
type MockRepository[T any] struct {
	CreateFunc            func(entity *T) *models.ServiceError
	GetByIDFunc           func(id int64, preloads ...string) (*T, *models.ServiceError)
	UpdateFunc            func(entity *T) *models.ServiceError
	UpdateFieldsFunc      func(id int64, updates map[string]interface{}) (*T, *models.ServiceError)
	DeleteFunc            func(id int64) *models.ServiceError
	DeleteEntityFunc      func(entity *T) *models.ServiceError
	ListFunc              func(offset, limit int, preloads ...string) (*[]T, *models.ServiceError)
	CountFunc             func() (*int64, *models.ServiceError)
	FindByFieldFunc       func(field string, value interface{}, preloads ...string) (*T, *models.ServiceError)
	FindAllByFieldFunc    func(field string, value interface{}, offset, limit int, preloads ...string) (*[]T, *models.ServiceError)

	// Track method calls
	CreateCalled       bool
	GetByIDCalled      bool
	UpdateCalled       bool
	UpdateFieldsCalled bool
	DeleteCalled       bool
	DeleteEntityCalled bool
	ListCalled         bool
	CountCalled        bool
	FindByFieldCalled  bool
	FindAllByFieldCalled bool

	// Store method arguments
	LastCreateEntity         *T
	LastGetByIDId           int64
	LastGetByIDPreloads     []string
	LastUpdateEntity        *T
	LastUpdateFieldsId      int64
	LastUpdateFieldsUpdates map[string]interface{}
	LastDeleteId            int64
	LastDeleteEntity        *T
	LastListOffset          int
	LastListLimit           int
	LastListPreloads        []string
	LastFindByFieldField    string
	LastFindByFieldValue    interface{}
	LastFindByFieldPreloads []string
}

// NewMockRepository creates a new mock repository
func NewMockRepository[T any]() *MockRepository[T] {
	return &MockRepository[T]{}
}

func (m *MockRepository[T]) Create(entity *T) *models.ServiceError {
	m.CreateCalled = true
	m.LastCreateEntity = entity
	if m.CreateFunc != nil {
		return m.CreateFunc(entity)
	}
	return nil
}

func (m *MockRepository[T]) GetByID(id int64, preloads ...string) (*T, *models.ServiceError) {
	m.GetByIDCalled = true
	m.LastGetByIDId = id
	m.LastGetByIDPreloads = preloads
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id, preloads...)
	}
	return nil, nil
}

func (m *MockRepository[T]) Update(entity *T) *models.ServiceError {
	m.UpdateCalled = true
	m.LastUpdateEntity = entity
	if m.UpdateFunc != nil {
		return m.UpdateFunc(entity)
	}
	return nil
}

func (m *MockRepository[T]) UpdateFields(id int64, updates map[string]interface{}) (*T, *models.ServiceError) {
	m.UpdateFieldsCalled = true
	m.LastUpdateFieldsId = id
	m.LastUpdateFieldsUpdates = updates
	if m.UpdateFieldsFunc != nil {
		return m.UpdateFieldsFunc(id, updates)
	}
	return nil, nil
}

func (m *MockRepository[T]) Delete(id int64) *models.ServiceError {
	m.DeleteCalled = true
	m.LastDeleteId = id
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func (m *MockRepository[T]) DeleteEntity(entity *T) *models.ServiceError {
	m.DeleteEntityCalled = true
	m.LastDeleteEntity = entity
	if m.DeleteEntityFunc != nil {
		return m.DeleteEntityFunc(entity)
	}
	return nil
}

func (m *MockRepository[T]) List(offset, limit int, preloads ...string) (*[]T, *models.ServiceError) {
	m.ListCalled = true
	m.LastListOffset = offset
	m.LastListLimit = limit
	m.LastListPreloads = preloads
	if m.ListFunc != nil {
		return m.ListFunc(offset, limit, preloads...)
	}
	return nil, nil
}

func (m *MockRepository[T]) Count() (*int64, *models.ServiceError) {
	m.CountCalled = true
	if m.CountFunc != nil {
		return m.CountFunc()
	}
	count := int64(0)
	return &count, nil
}

func (m *MockRepository[T]) FindByField(field string, value interface{}, preloads ...string) (*T, *models.ServiceError) {
	m.FindByFieldCalled = true
	m.LastFindByFieldField = field
	m.LastFindByFieldValue = value
	m.LastFindByFieldPreloads = preloads
	if m.FindByFieldFunc != nil {
		return m.FindByFieldFunc(field, value, preloads...)
	}
	return nil, nil
}

func (m *MockRepository[T]) FindAllByField(field string, value interface{}, offset, limit int, preloads ...string) (*[]T, *models.ServiceError) {
	m.FindAllByFieldCalled = true
	if m.FindAllByFieldFunc != nil {
		return m.FindAllByFieldFunc(field, value, offset, limit, preloads...)
	}
	return nil, nil
}

func (m *MockRepository[T]) Query() interface{} {
	return nil
}

// Reset clears all tracking data
func (m *MockRepository[T]) Reset() {
	m.CreateCalled = false
	m.GetByIDCalled = false
	m.UpdateCalled = false
	m.UpdateFieldsCalled = false
	m.DeleteCalled = false
	m.DeleteEntityCalled = false
	m.ListCalled = false
	m.CountCalled = false
	m.FindByFieldCalled = false
	m.FindAllByFieldCalled = false
	
	m.LastCreateEntity = nil
	m.LastGetByIDId = 0
	m.LastGetByIDPreloads = nil
	m.LastUpdateEntity = nil
	m.LastUpdateFieldsId = 0
	m.LastUpdateFieldsUpdates = nil
	m.LastDeleteId = 0
	m.LastDeleteEntity = nil
	m.LastListOffset = 0
	m.LastListLimit = 0
	m.LastListPreloads = nil
	m.LastFindByFieldField = ""
	m.LastFindByFieldValue = nil
	m.LastFindByFieldPreloads = nil
}