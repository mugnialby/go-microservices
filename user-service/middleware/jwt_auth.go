package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mugnialby/go-microservices/user-service/models/dto/responses"
	"github.com/spf13/viper"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerToken := context.GetHeader("Authorization")
		if headerToken == "" {
			context.JSON(http.StatusUnauthorized, responses.NewResponse("Unauthorized", nil))
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
			context.JSON(http.StatusUnauthorized, responses.NewResponse("Unauthorized", nil))
			context.Abort()
			return
		}

		context.Next()
	}
}
