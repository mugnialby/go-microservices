package servicesimpl

import (
	"github.com/go-playground/validator"
	"github.com/mugnialby/go-microservices/auth-service/services"
)

type validationServiceImpl struct {
	validate *validator.Validate
}

func NewValidationService() services.ValidationService {
	return &validationServiceImpl{
		validate: validator.New(),
	}
}

func (validationService *validationServiceImpl) Validate(data interface{}) error {
	return validationService.validate.Struct(data)
}
