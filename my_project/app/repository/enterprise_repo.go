package repository

import (
	"gorm.io/gorm"
	"your-project/app/domain"
)

type EnterpriseRepository struct {
	DB *gorm.DB
}

func (r *EnterpriseRepository) CreateEnterprise(enterprise *domain.Enterprise) error {
	return r.DB.Create(enterprise).Error
}
