package servicesimplmocks

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type ValidationServiceMock struct {
	Mock mock.Mock
}

func (validationServiceMock *ValidationServiceMock) Validate(data interface{}) error {
	arguments := validationServiceMock.Mock.Called(data)

	if arguments.Get(0) == nil {
		return errors.New("PARAMETERS IS NULL")
	} else {
		return nil
	}
}
