//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mugnialby/go-microservices/auth-service/connections"
	"github.com/mugnialby/go-microservices/auth-service/handlers"
	proto "github.com/mugnialby/go-microservices/auth-service/proto/users"
	"github.com/mugnialby/go-microservices/auth-service/services"
	servicesimpl "github.com/mugnialby/go-microservices/auth-service/services/impl"
)

/**
 * created by	: albym
 * created at	: 5 Jun 2024
 *
 * this class will generate providers for dependency injection
 */
func ProvideAuthService() services.AuthService {
	return servicesimpl.NewAuthService()
}

func ProvideValidationService() services.ValidationService {
	return servicesimpl.NewValidationService()
}

func ProvideUserGrpcConnection() proto.UserServiceClient {
	return connections.NewUserConnection()
}

func ProvideAuthHandler(
	authService services.AuthService,
	validationService services.ValidationService,
	userServiceClient proto.UserServiceClient,
) *handlers.AuthHandler {
	return handlers.NewAuthHandler(authService, validationService, userServiceClient)
}

func InitAuthHandler() *handlers.AuthHandler {
	wire.Build(ProvideAuthService, ProvideAuthHandler, ProvideValidationService, ProvideUserGrpcConnection)
	return &handlers.AuthHandler{}
}
