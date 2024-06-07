package repositories

import "github.com/mugnialby/go-microservices/user-service/models/entities"

type UserRepository interface {
	FindAll() (users *[]entities.Users, err error)
	FindById(userId *uint64) (user *entities.Users, err error)
	FindByUsername(username string) (user *entities.Users, err error)
	Add(user *entities.Users) (err error)
	Update(user *entities.Users) (err error)
	Delete(userId *uint64) (err error)
}
