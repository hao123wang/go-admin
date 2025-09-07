package middleware

import (
	"go-admin-server/api/dao"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"go-admin-server/global"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var logDao = dao.SysLogDao{}

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		userInfo, exists := c.Get(global.LoggedUser)
		if !exists {
			response.Error(c, response.ErrAdminUnauthorized)
			c.Abort()
			return
		}
		loggedUser, ok := userInfo.(entity.JwtAdmin)
		if !ok {
			global.Logger.Error("Failed to assert operation user")
			response.Error(c, response.ErrServerError)
			c.Abort()
			return
		}

		operationLog := &entity.SysOperationLog{
			AdminID:   loggedUser.ID,
			Username:  loggedUser.Username,
			Method:    method,
			Ip:        c.ClientIP(),
			Url:       c.Request.URL.Path,
			CreatedAt: utils.HTime{Time: time.Now()},
		}

		if err := logDao.CreateOperationLog(operationLog); err != nil {
			response.Error(c, response.ErrServerError)
			c.Abort()
			return
		}

		c.Next()
	}
}
