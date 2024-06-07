package services

import (
	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/user-service/models/entities"
)

type UserService interface {
	FindAll() (users *[]entities.Users, err error)
	FindById(userId *uint64) (user *entities.Users, err error)
	FindByUsername(username string) (user *entities.Users, err error)
	Add(userAddRequest *requests.UsersAddRequest) (err error)
	Update(userId *uint64, userAddRequest *requests.UsersUpdateRequest) (err error)
	Delete(userId *uint64) (err error)
}
