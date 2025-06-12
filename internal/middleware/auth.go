package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tc-api/internal/configs"
)

func InitAuth(config *configs.Configuration) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := allowedSecretKey(config, c.GetHeader("Authorization")); err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

func allowedSecretKey(config *configs.Configuration, requestHeader string) (err error) {
	if requestHeader == "" {
		err = fmt.Errorf("no authorization header found")
		return
	}
	var scheme, secretKey string
	_, err = fmt.Sscanf(requestHeader, "%s %s", &scheme, &secretKey)
	if err != nil {
		err = fmt.Errorf("invalid authorization Format")
		return
	}
	if scheme != "Bearer" {
		err = fmt.Errorf("invalid authorization Format")
		return
	}
	if secretKey != config.App.SecretKey {
		err = fmt.Errorf("invalid secret key")
		return
	}
	return
}
