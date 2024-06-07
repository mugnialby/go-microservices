package services

type ValidationService interface {
	Validate(interface{}) error
}
