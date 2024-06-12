package servicesimpltests

import (
	"testing"

	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	servicesimpl "github.com/mugnialby/go-microservices/user-service/services/impl"
	"github.com/stretchr/testify/assert"
)

var validationService = servicesimpl.NewValidationService()

func TestValidate(t *testing.T) {
	userUpdateRequest := requests.UsersUpdateRequest{
		FirstName:  "TES",
		LastName:   "TES",
		ModifiedBy: "TES",
	}

	err := validationService.Validate(&userUpdateRequest)
	assert.Nil(t, err)
}

func TestValidateFail(t *testing.T) {
	userUpdateRequest := requests.UsersUpdateRequest{
		FirstName:  "",
		LastName:   "TES",
		ModifiedBy: "TES",
	}

	err := validationService.Validate(&userUpdateRequest)
	assert.NotNil(t, err)
}
