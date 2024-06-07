package repositoriesimpl

import (
	"errors"

	"github.com/mugnialby/go-microservices/user-service/models/entities"
	"github.com/mugnialby/go-microservices/user-service/repositories"
	"github.com/mugnialby/go-microservices/user-service/utils"
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

func (userRepository *userRepositoryImpl) FindAll() (users *[]entities.Users, err *error) {
	results := userRepository.db.Find(&users)
	utils.ErrorHandler(results.Error)

	return users, nil
}

func (userRepository *userRepositoryImpl) FindById(userId *uint64) (user *entities.Users, err *error) {
	result := userRepository.db.First(&user, userId)
	utils.ErrorHandler(result.Error)

	if result.RowsAffected == 0 {
		errMessage := errors.New("User is not found")
		return nil, &errMessage
	}

	return user, nil
}

func (userRepository *userRepositoryImpl) Add(user *entities.Users) (err *error) {
	result := userRepository.db.Create(&user)
	utils.ErrorHandler(result.Error)

	return nil
}

func (userRepository *userRepositoryImpl) Update(user *entities.Users) (err *error) {
	result := userRepository.db.Model(&user).Updates(user)
	utils.ErrorHandler(result.Error)

	return nil
}

func (userRepository *userRepositoryImpl) Delete(userId *uint64) (err *error) {
	var user entities.Users

	result := userRepository.db.Where("id = ?", userId).Delete(&user)
	utils.ErrorHandler(result.Error)

	return nil
}
