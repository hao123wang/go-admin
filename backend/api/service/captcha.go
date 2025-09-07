package service

import (
	"go-admin-server/api/dao"
	"go-admin-server/global"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaStore = &dao.CaptcahStore{}

// 生成验证码
func CaptchaMake() (string, string) {

	// 配置验证码信息
	driver := base64Captcha.DriverString{
		Height:     80,
		Width:      240,
		Length:     6,
		Source:     "1234567890qwertyuiopasdfghjklzxcvbnm",
		NoiseCount: 0,
		Fonts:      []string{"wqy-microhei.ttc"},
	}

	cp := base64Captcha.NewCaptcha(&driver, captchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.Logger.Error("Failed to generate captcha", zap.Error(err))
		return "", ""
	}
	return id, b64s
}
