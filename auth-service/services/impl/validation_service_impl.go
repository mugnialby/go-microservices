package servicesimpl

import "github.com/go-playground/validator"

type validationServiceImpl struct {
	validate *validator.Validate
}

func NewValidationService() *validationServiceImpl {
	return &validationServiceImpl{
		validate: validator.New(),
	}
}

func (validationService *validationServiceImpl) Validate(data interface{}) error {
	return validationService.validate.Struct(data)
}
