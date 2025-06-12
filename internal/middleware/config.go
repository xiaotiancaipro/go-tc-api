package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tc-api/internal/configs"
)

func InitConfig(config *configs.Configuration) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func GetConfig(c *gin.Context) (config *configs.Configuration, err error) {
	cConfig, exists := c.Get("config")
	if !exists {
		err = fmt.Errorf("failed to get config context")
		return
	}
	gConfig, ok := cConfig.(*configs.Configuration)
	if !ok {
		err = fmt.Errorf("type assertion failed")
		return
	}
	config = gConfig
	return
}
