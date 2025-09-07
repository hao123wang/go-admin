package core

import (
	"context"
	"fmt"
	"go-admin-server/global"
	"go-admin-server/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func RunServer() {
	router := router.SetupRouter()
	address := fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port)

	// 配置服务器
	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	// 优雅启动
	go func() {
		global.Logger.Info("Starting server", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	// 设置优雅关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// 优雅关闭
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}
	global.Logger.Info("Server exit")
}
