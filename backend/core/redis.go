// 初始化 Redis

package core

import (
	"context"
	"go-admin-server/global"

	"github.com/redis/go-redis/v9"
)

func InitRDB() *redis.Client {
	config := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return client
}
