package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUtil struct{}

func (r ResponseUtil) ParsingRequest(c *gin.Context, data any) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		r.BadRequest(c)
		return false
	}
	return true
}

func (ResponseUtil) Build(c *gin.Context, code int, message string, data any) {
	c.JSON(code, gin.H{"code": code, "message": message, "data": data})
}

func (r ResponseUtil) Unauthorized(c *gin.Context, message string) {
	r.Build(c, http.StatusUnauthorized, "Error: "+message, nil)
}

func (r ResponseUtil) BadRequest(c *gin.Context) {
	r.Build(c, http.StatusBadRequest, "Error: 请求参数解析错误", nil)
}

func (r ResponseUtil) Success(c *gin.Context, message string, data any) {
	r.Build(c, http.StatusOK, "Success: "+message, data)
}

func (r ResponseUtil) SuccessNoMessage(c *gin.Context, data any) {
	r.Build(c, http.StatusOK, "Success", data)
}

func (r ResponseUtil) Error(c *gin.Context, message string) {
	r.Build(c, http.StatusInternalServerError, "Error: 内部错误, "+message, nil)
}
