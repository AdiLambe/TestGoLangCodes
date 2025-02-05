package repository

import (
	"gorm.io/gorm"
	"your-project/app/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}
