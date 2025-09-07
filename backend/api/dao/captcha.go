package dao

import (
	"context"
	"go-admin-server/common/response"
	"go-admin-server/global"
	"time"

	"go.uber.org/zap"
)

var ctx = context.Background()

type CaptcahStore struct{} // 需要实现 base64Captcha.Store 接口

// Set 实现 Set 方法，存储验证码
func (store *CaptcahStore) Set(id, value string) error {
	key := global.CaptchaPrex + id
	err := global.RDB.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		global.Logger.Error("Failed to set captcha", zap.Error(err))
		return response.ErrServerError
	}
	return nil
}

// Get 实现 Get 方法，获取验证码
func (store *CaptcahStore) Get(id string, clear bool) string {
	key := global.CaptchaPrex + id
	value, err := global.RDB.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		go global.RDB.Del(ctx, key)
	}
	return value
}

// 验证
func (store *CaptcahStore) Verify(id, answer string, clear bool) bool {
	storedValue := store.Get(id, clear)
	return storedValue == answer
}
