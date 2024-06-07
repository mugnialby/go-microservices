package main

import (
	"github.com/mugnialby/go-microservices/user-service/config"
	"github.com/mugnialby/go-microservices/user-service/server"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * main project file
 * will get config data in /config/config.yaml
 * will initiate database connection
 * will initiate dependency injection
 */
func main() {
	config.InitConfig()
	db := config.InitDbConnection()

	userHandler := InitUserHandler(db)
	userGrpcService := InitUserGrpcService(db)

	go func() {
		server.StartGrpcServer(userGrpcService)
	}()

	server.StartHttpServer(userHandler)

	select {}
}
