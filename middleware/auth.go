package middleware

import (
	"server/core/response"

	"github.com/gin-gonic/gin"
)

// TokenAuth Token认证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token是否为空
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Fail(c, response.TokenEmpty)
			c.Abort()
			return
		}
	}
}
