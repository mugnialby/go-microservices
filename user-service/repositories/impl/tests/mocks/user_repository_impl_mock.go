package repositoriesmock

import (
	"errors"
	"time"

	"github.com/mugnialby/go-microservices/user-service/models/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepositoryMock *UserRepositoryMock) FindAll() (users *[]entities.Users, err error) {
	arguments := userRepositoryMock.Mock.Called()

	if arguments.Get(0) == nil {
		return nil, errors.New("DATA NOT FOUND")
	}

	var results []entities.Users

	result := entities.Users{
		ID:         1,
		Username:   "ALBYM",
		Password:   "12345678",
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		Email:      "ALBY@TES.COM",
		ManagerId:  2,
		Status:     "Y",
		CreatedBy:  "ALBYM",
		CreatedAt:  time.Now(),
		ModifiedBy: "",
		ModifiedAt: time.Time{},
	}

	results = append(results, result)

	return &results, nil
}

func (userRepositoryMock *UserRepositoryMock) FindById(userId *uint64) (user *entities.Users, err error) {
	arguments := userRepositoryMock.Mock.Called(userId)

	if arguments.Get(0) == nil {
		return nil, errors.New("DATA NOT FOUND")
	}

	result := entities.Users{
		ID:         *userId,
		Username:   "ALBYM",
		Password:   "12345678",
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		Email:      "ALBY@TES.COM",
		ManagerId:  2,
		Status:     "Y",
		CreatedBy:  "ALBYM",
		CreatedAt:  time.Now(),
		ModifiedBy: "",
		ModifiedAt: time.Time{},
	}

	return &result, nil
}

func (userRepositoryMock *UserRepositoryMock) Add(user *entities.Users) (err error) {
	arguments := userRepositoryMock.Mock.Called(user)

	if arguments.Get(0) == nil {
		return errors.New("PARAMETERS IS NULL")
	}

	return nil
}

func (userRepositoryMock *UserRepositoryMock) Update(user *entities.Users) (err error) {
	arguments := userRepositoryMock.Mock.Called(user)

	if arguments.Get(0) == nil {
		return errors.New("DATA NOT FOUND")
	} else {
		return nil
	}
}

func (userRepositoryMock *UserRepositoryMock) Delete(userId *uint64) (err error) {
	arguments := userRepositoryMock.Mock.Called(userId)

	if arguments.Get(0) == nil {
		return errors.New("DATA NOT FOUND")
	}

	return nil
}

func (userRepositoryMock *UserRepositoryMock) FindByUsername(username string) (user *entities.Users, err error) {
	arguments := userRepositoryMock.Mock.Called(username)

	if arguments.Get(0) == nil {
		return nil, errors.New("DATA NOT FOUND")
	}

	result := entities.Users{
		ID:         1,
		Username:   "ALBYM",
		Password:   "12345678",
		FirstName:  "ALBY",
		LastName:   "MUGNI",
		Email:      "ALBY@TES.COM",
		ManagerId:  2,
		Status:     "Y",
		CreatedBy:  "ALBYM",
		CreatedAt:  time.Now(),
		ModifiedBy: "",
		ModifiedAt: time.Time{},
	}

	return &result, nil
}
