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
func ProvideJwtService() services.JwtService {
	return servicesimpl.NewJwtService()
}

func ProvideValidationService() services.ValidationService {
	return servicesimpl.NewValidationService()
}

func ProvideUserGrpcConnection() proto.UserServiceClient {
	return connections.NewUserConnection()
}

func ProvideAuthHandler(
	validationService services.ValidationService,
	userServiceClient proto.UserServiceClient,
	jwtService services.JwtService,
) *handlers.AuthHandler {
	return handlers.NewAuthHandler(validationService, userServiceClient, jwtService)
}

func InitAuthHandler() *handlers.AuthHandler {
	wire.Build(ProvideAuthHandler, ProvideValidationService, ProvideUserGrpcConnection, ProvideJwtService)
	return &handlers.AuthHandler{}
}
