package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/user-service/handlers"
	"github.com/mugnialby/go-microservices/user-service/routes"
	"github.com/spf13/viper"
)

func StartHttpServer(
	userHandler *handlers.UserHandler,
) {
	router := gin.Default()

	// router.Use(middleware.JWTAuth())
	routes.InitUserRoutes(router, userHandler)

	router.Run(viper.GetString("server.port"))
}
