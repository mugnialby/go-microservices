package config

import (
	"fmt"
	"strings"

	"github.com/mugnialby/go-microservices/user-service/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * initiate database connection for gorm
 * @return db (pointer to gorm.DB)
 */
func InitDbConnection() *gorm.DB {
	var dbConnection strings.Builder
	fmt.Fprintf(&dbConnection, "host=%s", viper.GetString("database.host"))
	fmt.Fprintf(&dbConnection, " user=%s", viper.GetString("database.user"))
	fmt.Fprintf(&dbConnection, " password=%s", viper.GetString("database.password"))
	fmt.Fprintf(&dbConnection, " dbname=%s", viper.GetString("database.dbname"))
	fmt.Fprintf(&dbConnection, " port=%s", viper.GetString("database.port"))
	fmt.Fprintf(&dbConnection, " sslmode=%s", viper.GetString("database.sslmode"))
	fmt.Fprintf(&dbConnection, " TimeZone=%s", viper.GetString("database.TimeZone"))

	db, err := gorm.Open(postgres.Open(dbConnection.String()), &gorm.Config{})
	utils.ErrorHandler(err)

	return db
}
