package dto

type CreateUserRequest struct {
	UserName     string `json:"userName" validate:"required,min=3,max=20"`
	Email        string `json:"email" validate:"required,email"`
	EnterpriseID string `json:"enterpriseId" validate:"required,uuid"`
}
