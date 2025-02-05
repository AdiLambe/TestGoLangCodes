package repository

import (
	"errors"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/domain"
	"sync"

	"github.com/google/uuid"
)

// EnterpriseRepository is an in-memory store for enterprises
type EnterpriseRepository struct {
	data map[uuid.UUID]*domain.Enterprise
	mu   sync.RWMutex
}

// NewEnterpriseRepository initializes the repository
func NewEnterpriseRepository() *EnterpriseRepository {
	return &EnterpriseRepository{
		data: make(map[uuid.UUID]*domain.Enterprise),
	}
}

// CreateEnterprise adds a new enterprise
func (r *EnterpriseRepository) CreateEnterprise(enterprise *domain.Enterprise) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	enterprise.ID = uuid.New()
	r.data[enterprise.ID] = enterprise
	return nil
}

// GetAllEnterprises retrieves all enterprises
func (r *EnterpriseRepository) GetAllEnterprises() ([]domain.Enterprise, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var enterprises []domain.Enterprise
	for _, enterprise := range r.data {
		enterprises = append(enterprises, *enterprise)
	}
	return enterprises, nil
}
