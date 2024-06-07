package servicesimpl

import "github.com/mugnialby/go-microservices/auth-service/services"

type authServiceImpl struct {
}

func NewAuthService() services.AuthService {
	return &authServiceImpl{}
}

func (authService *authServiceImpl) Login() {

}
