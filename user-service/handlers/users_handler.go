package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/user-service/models/dto/requests"
	"github.com/mugnialby/go-microservices/user-service/models/dto/responses"
	"github.com/mugnialby/go-microservices/user-service/services"
	"github.com/sirupsen/logrus"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * Rest Controller (for handling requests)
 */
type UserHandler struct {
	userService       services.UserService
	validationService services.ValidationService
}

func NewUserHandler(userService services.UserService, validationService services.ValidationService) *UserHandler {
	return &UserHandler{
		userService:       userService,
		validationService: validationService,
	}
}

func (handler *UserHandler) UserFindAllHandler(context *gin.Context) {
	users, err := handler.userService.FindAll()
	if err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusInternalServerError, responses.NewResponse("Fail", nil))
		return
	}

	context.JSON(http.StatusOK, responses.NewResponse("OK", &users))
}

func (handler *UserHandler) UserFindByIdHandler(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 0, 64)
	if err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	user, err := handler.userService.FindById(&userId)
	if err != nil {
		if err.Error() == "record not found" {
			logrus.Printf("Record not found | Parameter id: %v", &userId)
			context.JSON(http.StatusNotFound, responses.NewResponse("Data not found", nil))
		} else {
			logrus.Errorf("Internal Server Error : %v", err)
			context.JSON(http.StatusInternalServerError, responses.NewResponse(err.Error(), nil))
		}

		return
	}

	context.JSON(http.StatusOK, responses.NewResponse("OK", &user))
}

func (handler *UserHandler) UserAddHandler(context *gin.Context) {
	var userAddRequest requests.UsersAddRequest
	err := context.BindJSON(&userAddRequest)
	if err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorValidation := handler.validationService.Validate(&userAddRequest)
	if errorValidation != nil {
		logrus.Printf("%v | Parameter : %v", errorValidation.Error(), &userAddRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorService := handler.userService.Add(&userAddRequest)
	if errorService != nil {
		logrus.Errorln(errorService.Error())
		context.JSON(http.StatusInternalServerError, responses.NewResponse("Fail", nil))
		return
	}

	context.JSON(http.StatusCreated, responses.NewResponse("OK", nil))
}

func (handler *UserHandler) UserUpdateHandler(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 0, 64)
	if err != nil {
		logrus.Errorln(err.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	_, errorFindData := handler.userService.FindById(&userId)
	if errorFindData != nil {
		if errorFindData.Error() == "record not found" {
			logrus.Printf("Record not found | Parameter id: %v", &userId)
			context.JSON(http.StatusNotFound, responses.NewResponse("Data not found", nil))
		} else {
			logrus.Errorf("Error : %v | Parameter : %v", errorFindData, &userId)
			context.JSON(http.StatusInternalServerError, responses.NewResponse(errorFindData.Error(), nil))
		}

		return
	}

	var userUpdateRequest requests.UsersUpdateRequest
	errorBindingJson := context.BindJSON(&userUpdateRequest)
	if errorBindingJson != nil {
		logrus.Printf("Error : %v | Parameter : %v", errorBindingJson, &userUpdateRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorValidation := handler.validationService.Validate(&userUpdateRequest)
	if errorValidation != nil {
		logrus.Printf("Error : %v | Parameter : %v", errorValidation.Error(), &userUpdateRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	errorUpdate := handler.userService.Update(&userId, &userUpdateRequest)
	if errorUpdate != nil {
		logrus.Errorf("Error : %v | ID: %v | Parameter : %v", errorUpdate.Error(), &userId, &userUpdateRequest)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Fail", nil))
		return
	}

	context.JSON(http.StatusOK, responses.NewResponse("OK", nil))
}

func (handler *UserHandler) UserDeleteHandler(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 0, 64)
	if err != nil {
		logrus.Printf(err.Error())
		context.JSON(http.StatusBadRequest, responses.NewResponse("Bad Request", nil))
		return
	}

	_, errorFindData := handler.userService.FindById(&userId)
	if errorFindData != nil {
		if errorFindData.Error() == "record not found" {
			logrus.Printf("Record Not Found | Parameter ID : %v", &userId)
			context.JSON(http.StatusNotFound, responses.NewResponse("Data not found", nil))
		} else {
			logrus.Errorf("Error : %v", errorFindData)
			context.JSON(http.StatusInternalServerError, responses.NewResponse(errorFindData.Error(), nil))
		}

		return
	}

	errorDeleteData := handler.userService.Delete(&userId)
	if errorDeleteData != nil {
		logrus.Errorf("Error : %v | Parameter ID : %v", errorDeleteData.Error(), &userId)
		context.JSON(http.StatusBadRequest, responses.NewResponse("Fail", nil))
		return
	}

	context.JSON(http.StatusOK, responses.NewResponse("OK", nil))
}
