package middleware

import (
	"go-tc-api/internal/configs"
	"go-tc-api/internal/utils"
)

var response = utils.ResponseUtil{}

type Middleware struct {
	Config *configs.Configuration
}
