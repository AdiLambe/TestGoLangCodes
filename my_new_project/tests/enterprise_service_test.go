package tests

import (
	"testing"

	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/repository"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/service"

	"github.com/stretchr/testify/assert"
)

func TestCreateEnterprise(t *testing.T) {
	repo := repository.NewEnterpriseRepository()
	service := service.NewEnterpriseService(repo)

	enterprise := &domain.Enterprise{Name: "TestCorp", Region: "us-east-1"}
	err := service.CreateEnterprise(enterprise)

	assert.NoError(t, err)
}

func TestGetAllEnterprises(t *testing.T) {
	repo := repository.NewEnterpriseRepository()
	service := service.NewEnterpriseService(repo)

	service.CreateEnterprise(&domain.Enterprise{Name: "TestCorp", Region: "us-east-1"})
	enterpriseList, err := service.GetAllEnterprises()

	assert.NoError(t, err)
	assert.Len(t, enterpriseList, 1)
}
