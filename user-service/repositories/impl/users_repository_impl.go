package repositoriesimpl

import (
	"github.com/mugnialby/go-microservices/user-service/models/entities"
	"github.com/mugnialby/go-microservices/user-service/repositories"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (userRepository *userRepositoryImpl) FindAll() (users *[]entities.Users, err error) {
	results := userRepository.db.Find(&users)
	return users, results.Error
}

func (userRepository *userRepositoryImpl) FindById(userId *uint64) (user *entities.Users, err error) {
	result := userRepository.db.First(&user, userId)
	return user, result.Error
}

func (userRepository *userRepositoryImpl) Add(user *entities.Users) (err error) {
	result := userRepository.db.Create(&user)
	return result.Error
}

func (userRepository *userRepositoryImpl) Update(user *entities.Users) (err error) {
	result := userRepository.db.Model(&user).Updates(user)
	return result.Error
}

func (userRepository *userRepositoryImpl) Delete(userId *uint64) (err error) {
	result := userRepository.db.Delete(&entities.Users{}, userId)
	return result.Error
}

func (userRepository *userRepositoryImpl) FindByUsername(username string) (user *entities.Users, err error) {
	result := userRepository.db.Where("username = ?", username).First(&user)
	return user, result.Error
}
