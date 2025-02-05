package dto

type CreateEnterpriseRequest struct {
	Name   string `json:"name" validate:"required,min=10,max=100"`
	Region string `json:"region" validate:"required,oneof=us-east-1 us-east-2 asia-region-1"`
}
