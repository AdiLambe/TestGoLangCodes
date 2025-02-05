package tests

import (
	"testing"

	"github.com/AdiLambe/TestGoLangCodes/my_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/repository"
	"github.com/AdiLambe/TestGoLangCodes/my_project/app/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)

	user := &domain.User{
		UserName:     "TestUser",
		Email:        "test@example.com",
		EnterpriseID: uuid.New(),
	}

	err := service.CreateUser(user)
	assert.NoError(t, err)
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)

	user := &domain.User{UserName: "TestUser", Email: "invalid-email", EnterpriseID: uuid.New()}
	err := service.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, "invalid email format", err.Error())
}
