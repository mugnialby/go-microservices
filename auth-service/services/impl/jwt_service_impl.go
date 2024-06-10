package servicesimpl

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mugnialby/go-microservices/auth-service/services"
	"github.com/spf13/viper"
)

type jwtServiceImpl struct {
}

func NewJwtService() services.JwtService {
	return &jwtServiceImpl{}
}

func (jwtServiceImpl *jwtServiceImpl) GenerateJWT(email *string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})

	return token.SignedString([]byte(viper.GetString("auth.jwt.secret")))
}
