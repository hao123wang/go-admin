package controller

import (
	"go-admin-server/api/service"
	"go-admin-server/common/response"

	"github.com/gin-gonic/gin"
)

// @Summary 获取验证码
// @Description 获取验证码
// @Tags 无需认证接口
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptchaMake()
	response.SuccessWithData(c, map[string]string{
		"captcahId":    id,
		"captchaImage": base64Image,
	})
}
