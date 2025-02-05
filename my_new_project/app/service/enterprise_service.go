package service

import (
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/repository"
)

// EnterpriseService struct
type EnterpriseService struct {
	Repo *repository.EnterpriseRepository
}

// NewEnterpriseService initializes EnterpriseService
func NewEnterpriseService(repo *repository.EnterpriseRepository) *EnterpriseService {
	return &EnterpriseService{Repo: repo}
}

// CreateEnterprise adds a new enterprise
func (s *EnterpriseService) CreateEnterprise(enterprise *domain.Enterprise) error {
	return s.Repo.CreateEnterprise(enterprise)
}

// GetAllEnterprises retrieves all enterprises
func (s *EnterpriseService) GetAllEnterprises() ([]domain.Enterprise, error) {
	return s.Repo.GetAllEnterprises()
}
