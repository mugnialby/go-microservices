package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	workingDirectory, err := os.Getwd()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workingDirectory)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalln(err.Error())
	}

	gin.SetMode(viper.GetString("apps.mode"))
}
