package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" validate:"required,uuid"`
	UserName     string    `gorm:"size:20;not null" validate:"required,min=3,max=20"`
	Email        string    `gorm:"size:100;not null;unique" validate:"required,email"`
	EnterpriseID uuid.UUID `gorm:"type:uuid;not null" validate:"required,uuid"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
