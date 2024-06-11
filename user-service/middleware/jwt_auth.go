package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mugnialby/go-microservices/user-service/models/dto/responses"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerToken := context.GetHeader("Authorization")
		if headerToken == "" {
			logrus.Printf("TOKEN NOT EMBEDDED")
			context.JSON(http.StatusUnauthorized, responses.NewResponse("Unauthorized", nil))
			context.Abort()
			return
		}

		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New("INVALID SIGN METHOD")
				logrus.Warnf("Error : %v | Algorithm : %v", err.Error(), token.Header["alg"])
				return nil, err
			}
			return []byte(viper.GetString("auth.jwt.secret")), nil
		})

		if err != nil || !token.Valid {
			logrus.Printf(err.Error())
			context.JSON(http.StatusUnauthorized, responses.NewResponse("Unauthorized", nil))
			context.Abort()
			return
		}

		context.Next()
	}
}
