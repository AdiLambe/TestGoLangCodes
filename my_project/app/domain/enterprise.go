package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enterprise struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" validate:"required,uuid"`
	Name   string    `gorm:"size:100;not null" validate:"required,min=10,max=100"`
	Region string    `gorm:"size:20;not null" validate:"required,oneof=us-east-1 us-east-2 asia-region-1"`
}

func (e *Enterprise) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return nil
}
