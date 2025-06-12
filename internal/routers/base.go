package routers

import (
	"fmt"
	"github.com/fatedier/golib/log"
	"github.com/gin-gonic/gin"
	"go-tc-api/internal/middleware"
	"go-tc-api/internal/utils"
)

const (
	Health = "/health"
)

var response = utils.ResponseUtil{}

func useMiddleware(controller func(*gin.Context, *middleware.Middleware)) func(*gin.Context) {
	return func(c *gin.Context) {
		config, err := middleware.GetConfig(c)
		if err != nil {
			log.Error(err)
			response.Error(c, fmt.Sprintf("%v", err))
			return
		}
		controller(c, &middleware.Middleware{Config: config})
	}
}
