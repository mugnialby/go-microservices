package servicesimpltests

import (
	"testing"
	"time"

	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/user-service/models/entities"
	repositoriesmock "github.com/mugnialby/go-microservices/user-service/repositories/impl/tests/mocks"
	servicesimpl "github.com/mugnialby/go-microservices/user-service/services/impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = &repositoriesmock.UserRepositoryMock{Mock: mock.Mock{}}
var userService = servicesimpl.NewUserService(userRepositoryMock)

func TestFindAll(t *testing.T) {
	userRepositoryMock.Mock.On("FindAll").Return(1)

	users, err := userService.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestFindAllNoDataFound(t *testing.T) {
	userRepositoryMock.Mock.On("FindAll").Return(nil)

	users, err := userService.FindAll()

	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestFindById(t *testing.T) {
	var userId uint64 = 1
	userRepositoryMock.Mock.On("FindById", &userId).Return(1)

	users, err := userService.FindById(&userId)

	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestFindByIdNoDataFound(t *testing.T) {
	var userId uint64 = 1
	userRepositoryMock.Mock.On("FindById", &userId).Return(nil)

	users, err := userService.FindById(&userId)

	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestAdd(t *testing.T) {
	users := &entities.Users{
		Username:  "ALBYM",
		Password:  "12345678",
		FirstName: "ALBY",
		LastName:  "MUGNI",
		Email:     "ALBY@MAIL.COM",
		ManagerId: 2,
		Status:    "Y",
		CreatedBy: "ADMIN",
	}
	userRepositoryMock.Mock.On("Add", users).Return(1)

	addRequest := requests.UsersAddRequest{
		Username:  "ALBYM",
		Password:  "12345678",
		FirstName: "ALBY",
		LastName:  "MUGNI",
		Email:     "ALBY@MAIL.COM",
		ManagerId: 2,
		Status:    "Y",
		CreatedBy: "ADMIN",
	}

	err := userService.Add(&addRequest)

	assert.Nil(t, err)
}

func TestAddFailed(t *testing.T) {
	userRepositoryMock.Mock.On("Add", &entities.Users{}).Return(nil)

	err := userService.Add(&requests.UsersAddRequest{})

	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	var userId uint64 = 1
	userRepositoryMock.Mock.On("Delete", &userId).Return(1)

	err := userService.Delete(&userId)

	assert.Nil(t, err)
}

func TestDeleteFailed(t *testing.T) {
	var userId uint64 = 1
	userRepositoryMock.Mock.On("Delete", &userId).Return(nil)

	err := userService.Delete(&userId)

	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	userId := uint64(1)
	timeNow := time.Now()
	users := &entities.Users{
		ID:         1,
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		ModifiedBy: "ADMIN",
		ModifiedAt: timeNow,
	}
	userRepositoryMock.Mock.On("Update", users).Return(1)

	updateRequest := requests.UsersUpdateRequest{
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		ModifiedBy: "ADMIN",
	}

	err := userService.Update(&userId, &updateRequest)

	assert.Nil(t, err)
}

func TestUpdateFailed(t *testing.T) {
	userId := uint64(1)
	timeNow := time.Now()
	users := entities.Users{
		ID:         1,
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		ModifiedBy: "ADMIN",
		ModifiedAt: timeNow,
	}
	userRepositoryMock.Mock.On("Update", &users).Return(nil)

	updateRequest := requests.UsersUpdateRequest{
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		ModifiedBy: "ADMIN",
	}

	err := userService.Update(&userId, &updateRequest)

	assert.NotNil(t, err)
}

func TestFindByUsername(t *testing.T) {
	userRepositoryMock.Mock.On("FindByUsername", "ALBYM").Return(1)

	users, err := userService.FindByUsername("ALBYM")

	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestFindByUsernameNotFound(t *testing.T) {
	userRepositoryMock.Mock.On("FindByUsername", "NOT FOUND").Return(nil)

	users, err := userService.FindByUsername("NOT FOUND")

	assert.NotNil(t, err)
	assert.Nil(t, users)
}
