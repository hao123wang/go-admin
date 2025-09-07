// 初始化 Logger

package core

import (
	"go-admin-server/global"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	config := global.Config.Logger

	encoder := getEncoder()
	writeSyncer := getWriteSyncer(config.Filename, config.MaxSize, config.MaxAge, config.MaxBackups)
	if config.IsConsolePrint {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}
	var zapcoreLevel zapcore.Level
	if err := zapcoreLevel.UnmarshalText([]byte(config.Level)); err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, writeSyncer, zapcoreLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer(filename string, maxSize, maxAge, maxBackups int) zapcore.WriteSyncer {
	lumberJackerLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}
	return zapcore.AddSync(lumberJackerLogger)
}
