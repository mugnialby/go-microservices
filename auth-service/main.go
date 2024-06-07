package main

import (
	"github.com/mugnialby/go-microservices/auth-service/config"
	"github.com/mugnialby/go-microservices/auth-service/server"
)

/**
 * created by	: albym
 * created at	: 5 Jun 2024
 *
 * main project file
 * will get config data in /config/config.yaml
 * will initiate database connection
 * will initiate dependency injection
 */
func main() {
	config.InitConfig()

	authHandler := InitAuthHandler()

	server.StartHttpServer(authHandler)
}
