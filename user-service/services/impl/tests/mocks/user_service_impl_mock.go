package servicesimplmocks

import (
	"time"

	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/user-service/models/entities"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock mock.Mock
}

func (svc *userServiceImpl) FindAll() (users *[]entities.Users, err error) {
	return svc.userRepository.FindAll()
}

func (svc *userServiceImpl) FindById(userId *uint64) (user *entities.Users, err error) {
	return svc.userRepository.FindById(userId)
}

func (svc *userServiceImpl) Add(userAddRequest *requests.UsersAddRequest) (err error) {
	addRequest := entities.Users{
		Username:  userAddRequest.Username,
		Password:  userAddRequest.Password,
		FirstName: userAddRequest.FirstName,
		LastName:  userAddRequest.LastName,
		Email:     userAddRequest.Email,
		ManagerId: userAddRequest.ManagerId,
		Status:    userAddRequest.Status,
		CreatedBy: userAddRequest.CreatedBy,
	}

	return svc.userRepository.Add(&addRequest)
}

func (svc *userServiceImpl) Update(userId *uint64, userUpdateRequest *requests.UsersUpdateRequest) (err error) {
	updateRequest := entities.Users{
		ID:         *userId,
		FirstName:  userUpdateRequest.FirstName,
		LastName:   userUpdateRequest.LastName,
		ModifiedBy: userUpdateRequest.ModifiedBy,
		ModifiedAt: time.Now(),
	}

	return svc.userRepository.Update(&updateRequest)
}

func (svc *userServiceImpl) Delete(userId *uint64) (err error) {
	return svc.userRepository.Delete(userId)
}

func (svc *userServiceImpl) FindByUsername(username string) (user *entities.Users, err error) {
	return svc.userRepository.FindByUsername(username)
}
