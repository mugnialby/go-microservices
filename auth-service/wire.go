//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/mugnialby/go-microservices/auth-service/handlers"
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

func ProvideAuthHandler(authService services.AuthService) *handlers.AuthHandler {
	return handlers.NewAuthHandler(authService)
}

func ProvideValidationService() *validator.Validate {
	return servicesimpl.NewValidationService()
}

func InitAuthHandler() *handlers.AuthHandler {
	wire.Build(ProvideAuthService, ProvideAuthHandler, ProvideValidationService)
	return &handlers.AuthHandler{}
}
