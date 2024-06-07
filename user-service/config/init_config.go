package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/go-microservices/user-service/utils"
	"github.com/spf13/viper"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * initiate config file to load
 * will set mode for gin
 */
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		utils.ErrorHandler(err)
	}

	gin.SetMode(viper.GetString("apps.mode"))
}
