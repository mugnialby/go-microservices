package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerAuthenticationKey := context.GetHeader("X-API-KEY")
		if headerAuthenticationKey == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		authenticationKey := viper.GetString("auth.key")
		if headerAuthenticationKey != authenticationKey {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid"})
			context.Abort()
			return
		}

		context.Next()
	}
}
