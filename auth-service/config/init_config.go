package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("apps.mode"))
}
