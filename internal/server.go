package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tc-api/internal/configs"
	"go-tc-api/internal/middleware"
	"go-tc-api/internal/routers"
)

var (
	log    = configs.Logger()
	config configs.Configuration
	app    *gin.Engine
)

func GoTCAPIServer(configFile string) error {
	if err := configs.InitConfig(configFile, &config); err != nil {
		return err
	}
	initApp()
	initMiddleware()
	initRouter()
	addr := fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)
	log.Info("Go TC API Server is running on " + addr)
	if err := app.Run(addr); err != nil {
		return err
	}
	return nil
}

func initApp() {
	if config.App.Env != "production" {
		gin.SetMode(gin.DebugMode)
		debugInfo := "The current system environment is debug mode"
		log.Warning(debugInfo)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	app = gin.Default()
}

func initMiddleware() {
	app.Use(middleware.InitConfig(&config))
	app.Use(middleware.InitAuth(&config))
}

func initRouter() {
	routers.HealthRouter(app)
}
