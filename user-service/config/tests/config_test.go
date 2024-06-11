package configtest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mugnialby/go-microservices/user-service/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	wd, _ := os.Getwd()
	parent := filepath.Dir(wd)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(parent)

	err := viper.ReadInConfig()
	assert.Nil(t, err)
}

func TestDatabaseConnection(t *testing.T) {
	wd, _ := os.Getwd()
	parent := filepath.Dir(wd)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(parent)

	err := viper.ReadInConfig()
	db := config.InitDbConnection()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
