package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerToken := context.GetHeader("Authorization")
		if headerToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			context.Abort()
			return
		}

		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
			}
			return []byte(viper.GetString("auth.jwt.secret")), nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			context.Abort()
			return
		}

		context.Next()
	}
}

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(viper.GetString("auth.jwt.secret")))
}
