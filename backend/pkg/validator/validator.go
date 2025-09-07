package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// SetupValidator 初始化验证器
func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 使用 json 标签作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 注册自定义验证规则(可选内容)
		registerCustomValidations(v)

		// 注册自定义错误消息（可选内容）,用于覆盖默认的错误消息
	}
}

// registerCustomValidations 注册自定义验证规则
func registerCustomValidations(v *validator.Validate) {
	// 自定义手机号验证
	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		// 简单的手机号验证逻辑
		return len(phone) == 11 && strings.HasPrefix(phone, "1")
	})

	// 自定义密码强度验证
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		if len(password) < 8 {
			return false
		}
		// 至少包含字母和数字
		hasLetter := false
		hasDigit := false
		for _, char := range password {
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				hasLetter = true
			}
			if char >= '0' && char <= '9' {
				hasDigit = true
			}
		}
		return hasLetter && hasDigit
	})
}
