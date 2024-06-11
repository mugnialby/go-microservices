package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/responses"
	proto "github.com/mugnialby/go-microservices/auth-service/proto/users"
	"github.com/mugnialby/go-microservices/auth-service/services"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	validationService     services.ValidationService
	userClientGrpcService proto.UserServiceClient
	jwtService            services.JwtService
}

func NewAuthHandler(
	validationService services.ValidationService,
	userClientGrpcService proto.UserServiceClient,
	jwtService services.JwtService,
) *AuthHandler {
	return &AuthHandler{
		validationService:     validationService,
		userClientGrpcService: userClientGrpcService,
		jwtService:            jwtService,
	}
}

func (handler *AuthHandler) GenerateTokenHandler(context *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorValidation := handler.validationService.Validate(&loginRequest)
	if errorValidation != nil {
		logrus.Errorln(errorValidation.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	findByUsernameRequest := proto.FindByUsernameRequest{Username: loginRequest.Username}
	user, err := handler.userClientGrpcService.FindByUsername(context, &findByUsernameRequest)
	if err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusUnauthorized, responses.NewResponse("Bad Credential", nil))
		return
	}

	if user.Password != loginRequest.Password {
		context.JSON(http.StatusUnauthorized, responses.NewResponse("Bad Credential", nil))
		return
	}

	token, err := handler.jwtService.GenerateJWT(&user.Email)
	if err != nil {
		logrus.Warn(err.Error())
		context.JSON(http.StatusInternalServerError, responses.NewResponse("Fail", nil))
		return
	}

	context.JSON(http.StatusOK, responses.NewResponse("OK", token))
}
