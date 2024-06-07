package servicesimpl

import (
	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/user-service/models/entities"
	"github.com/mugnialby/go-microservices/user-service/repositories"
	"github.com/mugnialby/go-microservices/user-service/services"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) services.UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (svc *userServiceImpl) FindAll() (users *[]entities.Users, err *error) {
	return svc.userRepository.FindAll()
}

func (svc *userServiceImpl) FindById(userId *uint64) (user *entities.Users, err *error) {
	return svc.userRepository.FindById(userId)
}

func (svc *userServiceImpl) Add(userAddRequest *requests.UsersAddRequest) (err *error) {
	var addRequest = new(entities.Users)
	addRequest.Username = userAddRequest.Username
	addRequest.Password = userAddRequest.Password
	addRequest.FirstName = userAddRequest.FirstName
	addRequest.LastName = userAddRequest.LastName
	addRequest.Email = userAddRequest.Email
	addRequest.ManagerId = userAddRequest.ManagerId
	addRequest.Status = userAddRequest.Status
	addRequest.CreatedBy = userAddRequest.CreatedBy

	return svc.userRepository.Add(addRequest)
}

func (svc *userServiceImpl) Update(userUpdateRequest *requests.UsersUpdateRequest) (err *error) {
	var updateRequest = new(entities.Users)
	updateRequest.FirstName = userUpdateRequest.FirstName
	updateRequest.LastName = userUpdateRequest.LastName
	updateRequest.ModifiedBy = userUpdateRequest.ModifiedBy

	return svc.userRepository.Update(updateRequest)
}

func (svc *userServiceImpl) Delete(userId *uint64) (err *error) {
	return svc.userRepository.Delete(userId)
}
