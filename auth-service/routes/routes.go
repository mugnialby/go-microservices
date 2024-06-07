package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/auth-service/handlers"
)

func InitAuthRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
) {
	router.POST("auth/login", authHandler.LoginHandler)
}
