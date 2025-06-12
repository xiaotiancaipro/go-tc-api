package routers

import (
	"github.com/gin-gonic/gin"
	"go-tc-api/internal/controllers"
)

func HealthRouter(app *gin.Engine) {
	healthController := controllers.HealthController{}
	health := app.Group(Health)
	{
		health.GET("", healthController.Health)
	}
}
