package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/responses"
	proto "github.com/mugnialby/go-microservices/auth-service/proto/users"
	"github.com/mugnialby/go-microservices/auth-service/services"
)

type AuthHandler struct {
	authService           services.AuthService
	validationService     services.ValidationService
	userClientGrpcService proto.UserServiceClient
}

func NewAuthHandler(
	authService services.AuthService,
	validationService services.ValidationService,
	userClientGrpcService proto.UserServiceClient,
) *AuthHandler {
	return &AuthHandler{
		authService:           authService,
		validationService:     validationService,
		userClientGrpcService: userClientGrpcService,
	}
}

func (handler *AuthHandler) LoginHandler(context *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		log.Println(&loginRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorValidation := handler.validationService.Validate(&loginRequest)
	if errorValidation != nil {
		log.Println(errorValidation.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	findByUsernameRequest := proto.FindByUsernameRequest{Username: loginRequest.Username}
	user, err := handler.userClientGrpcService.FindByUsername(context, &findByUsernameRequest)
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusUnauthorized, responses.NewResponse("Bad Credential", nil))
		return
	}

	if user.Password != loginRequest.Password {
		context.JSON(http.StatusUnauthorized, responses.NewResponse("Bad Credential", nil))
		return
	}

	// token, err := middleware.GenerateJWT(*user.Email)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	// 	return
	// }

	// context.JSON(http.StatusOK, gin.H{"token": token})
	context.JSON(http.StatusOK, responses.NewResponse("token", nil))
}
