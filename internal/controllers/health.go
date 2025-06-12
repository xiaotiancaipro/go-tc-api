package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (HealthController) Health(c *gin.Context) {
	log.Info("Health check, App is running")
	response.SuccessNoMessage(c, nil)
}
