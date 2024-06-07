package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/auth-service/models/dto/responses"
	"github.com/mugnialby/go-microservices/auth-service/services"
)

type AuthHandler struct {
	authService       services.AuthService
	validationService *validator.Validate
}

func NewAuthHandler(authService services.AuthService, validationService *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authService:       authService,
		validationService: validationService,
	}
}

func (handler *AuthHandler) LoginHandler(context *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		log.Println(&loginRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	// user, err := handler.authService.FindById(&id)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, responses.NewResponse("Bad Credential", nil))
	// 	return
	// }

	// if user.Password != loginRequest.Password {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
	// 	return
	// }

	// token, err := middleware.GenerateJWT(*user.Email)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	// 	return
	// }

	// context.JSON(http.StatusOK, gin.H{"token": token})
	context.JSON(http.StatusOK, gin.H{"token": "tok"})
}
