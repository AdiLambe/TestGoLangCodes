package service

import (
	"github.com/google/uuid"
	"your-project/app/domain"
	"your-project/app/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(userName, email, enterpriseID string) (*domain.User, error) {
	enterpriseUUID, err := uuid.Parse(enterpriseID)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           uuid.New(),
		UserName:     userName,
		Email:        email,
		EnterpriseID: enterpriseUUID,
	}
	err = s.Repo.CreateUser(user)
	return user, err
}
