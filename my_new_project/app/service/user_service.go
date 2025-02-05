package service

import (
	"errors"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/repository"
	"regexp"
)

// UserService struct
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService initializes UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// CreateUser validates and adds a user
func (s *UserService) CreateUser(user *domain.User) error {
	if len(user.UserName) < 3 || len(user.UserName) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	return s.Repo.CreateUser(user)
}

// GetAllUsers retrieves all users of an enterprise
func (s *UserService) GetAllUsers(enterpriseID uuid.UUID) ([]domain.User, error) {
	return s.Repo.GetAllUsersByEnterpriseID(enterpriseID)
}

// Email validation helper
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
