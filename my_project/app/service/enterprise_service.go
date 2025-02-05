package service

import (
	"your-project/app/domain"
	"your-project/app/repository"

	"github.com/google/uuid"
)

type EnterpriseService struct {
	Repo *repository.EnterpriseRepository
}

func (s *EnterpriseService) CreateEnterprise(name, region string) (*domain.Enterprise, error) {
	enterprise := &domain.Enterprise{
		ID:     uuid.New(),
		Name:   name,
		Region: region,
	}
	err := s.Repo.CreateEnterprise(enterprise)
	return enterprise, err
}
