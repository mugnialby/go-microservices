package services

type JwtService interface {
	GenerateJWT(email *string) (token string, err error)
}
