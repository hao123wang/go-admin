package middleware

import (
	"go-admin-server/common/response"
	"go-admin-server/global"
	"go-admin-server/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization请求头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, response.ErrAdminUnauthorized)
			c.Abort()
			return
		}
		// 检查token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, response.ErrTokenFormatError)
			c.Abort()
			return
		}
		token := parts[1]

		// 解析token
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Error(c, response.ErrTokenInvalid)
			c.Abort()
			return
		}
		// 将当前登录用户的信息，设置到上下文中
		c.Set(global.LoggedUser, claims.JwtAdmin)
		c.Next()
	}
}
