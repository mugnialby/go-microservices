//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mugnialby/go-microservices/user-service/handlers"
	proto "github.com/mugnialby/go-microservices/user-service/proto/users"
	"github.com/mugnialby/go-microservices/user-service/repositories"
	repositoriesimpl "github.com/mugnialby/go-microservices/user-service/repositories/impl"
	"github.com/mugnialby/go-microservices/user-service/services"
	grpcserviceimpl "github.com/mugnialby/go-microservices/user-service/services/grpcimpl"
	servicesimpl "github.com/mugnialby/go-microservices/user-service/services/impl"
	"gorm.io/gorm"
)

/**
 * created by	: albym
 * created at	: 5 Jun 2024
 *
 * this class will generate providers for dependency injection
 */
func ProvideValidationService() services.ValidationService {
	return servicesimpl.NewValidationService()
}

func ProvideUserRepository(db *gorm.DB) repositories.UserRepository {
	return repositoriesimpl.NewUserRepository(db)
}

func ProvideUserService(userRepository repositories.UserRepository) services.UserService {
	return servicesimpl.NewUserService(userRepository)
}

func ProvideUserHandler(userService services.UserService, validationService services.ValidationService) *handlers.UserHandler {
	return handlers.NewUserHandler(userService, validationService)
}

func ProvideUserGrpcService(userService services.UserService) proto.UserServiceServer {
	return grpcserviceimpl.NewUserServiceServer(userService)
}

func InitUserHandler(db *gorm.DB) *handlers.UserHandler {
	wire.Build(ProvideUserRepository, ProvideUserService, ProvideUserHandler, ProvideValidationService)
	return &handlers.UserHandler{}
}

func InitUserService(db *gorm.DB) services.UserService {
	wire.Build(ProvideUserRepository, ProvideUserService)
	return nil
}

func InitUserGrpcService(db *gorm.DB) proto.UserServiceServer {
	wire.Build(ProvideUserGrpcService, ProvideUserService, ProvideUserRepository)
	return nil
}
