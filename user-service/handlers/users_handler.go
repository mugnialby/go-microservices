package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/user-service/services"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * Rest Controller (for handling requests)
 */
type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) UserFindAllHandler(context *gin.Context) {
	users, _ := handler.userService.FindAll()
	context.JSON(http.StatusOK, users)
}

func (handler *UserHandler) UserFindByIdHandler(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user, _ := handler.userService.FindById(&id)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *UserHandler) UserAddHandler(context *gin.Context) {
	// var userAddRequest requests.UsersAddRequest
	// err := context.BindJSON(&userAddRequest)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, responses.re)
	// }

	// err := handler.userService.Add(&userAddRequest)
	// if err != nil {
	// 	panic(err)
	// }

	context.JSON(http.StatusOK, "")
}

func (handler *UserHandler) UserUpdateHandler(context *gin.Context) {
	// id, err := strconv.ParseUint(context.Param("id"), 0, 64)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	// 	return
	// }

	// var userUpdateRequest requests.UsersUpdateRequest
	// utils.ErrorHandler(context.BindJSON(&userUpdateRequest))
	// err := handler.userService.Update(&userUpdateRequest)
	// if err != nil {
	// 	if err.Error() == "record not found" {
	// 		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	// 	} else {
	// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}

	// 	return
	// }

	// context.JSON(http.StatusOK, user)
	context.JSON(http.StatusOK, "")
}

func (handler *UserHandler) UserDeleteHandler(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user, _ := handler.userService.FindById(&id)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	context.JSON(http.StatusOK, user)
}

func TestHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "tes",
	})
}
