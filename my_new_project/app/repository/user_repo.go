package repository

import (
	"errors"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/domain"
	"sync"

	"github.com/google/uuid"
)

// UserRepository is an in-memory store for users
type UserRepository struct {
	data map[uuid.UUID]*domain.User
	mu   sync.RWMutex
}

// NewUserRepository initializes the repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: make(map[uuid.UUID]*domain.User),
	}
}

// CreateUser adds a new user
func (r *UserRepository) CreateUser(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.EnterpriseID == uuid.Nil {
		return errors.New("invalid enterprise ID")
	}

	user.ID = uuid.New()
	r.data[user.ID] = user
	return nil
}

// GetAllUsersByEnterpriseID fetches users for a given enterprise
func (r *UserRepository) GetAllUsersByEnterpriseID(enterpriseID uuid.UUID) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []domain.User
	for _, user := range r.data {
		if user.EnterpriseID == enterpriseID {
			users = append(users, *user)
		}
	}
	return users, nil
}
