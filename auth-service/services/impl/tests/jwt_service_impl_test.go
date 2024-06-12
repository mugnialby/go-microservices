package servicesimpltest

import (
	"testing"

	servicesimpl "github.com/mugnialby/go-microservices/auth-service/services/impl"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenSuccess(t *testing.T) {
	jwtService := servicesimpl.NewJwtService()

	email := "tes@mail.com"

	token, err := jwtService.GenerateJWT(&email)
	assert.NoError(t, err)
	assert.NotNil(t, token)
}
