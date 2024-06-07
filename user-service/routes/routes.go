package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/user-service/handlers"
)

func InitUserRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
) {
	router.GET("users/:id", userHandler.UserFindByIdHandler)
	router.GET("users/", userHandler.UserFindAllHandler)
	router.POST("users/", userHandler.UserAddHandler)
	router.PATCH("users/:id", userHandler.UserUpdateHandler)
	router.DELETE("users/:id", userHandler.UserDeleteHandler)
}
