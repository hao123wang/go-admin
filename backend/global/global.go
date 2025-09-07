package global

import (
	"go-admin-server/common/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.AppConfig
	Logger *zap.Logger
	DB     *gorm.DB
	RDB    *redis.Client
)
