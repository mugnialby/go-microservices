package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/auth-service/handlers"
	"github.com/mugnialby/go-microservices/auth-service/routes"
	"github.com/spf13/viper"
)

func StartHttpServer(
	authHandler *handlers.AuthHandler,
) {
	router := gin.Default()

	routes.InitAuthRoutes(router, authHandler)

	router.Run(viper.GetString("server.port"))
}
