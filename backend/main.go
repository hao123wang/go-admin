package main

import (
	"go-admin-server/common/config"
	"go-admin-server/common/flag"
	"go-admin-server/core"
	_ "go-admin-server/docs"
	"go-admin-server/global"
	"go-admin-server/pkg/validator"
)

// @title go-admin 后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT认证令牌，格式Bearer <token>
func main() {
	global.Config = config.Init()     // 配置文件
	global.Logger = core.InitLogger() // 日志
	global.DB = core.InitDB()         // MySQL
	global.RDB = core.InitRDB()       // Redis

	flag.InitFlag()            // 注册命令行工具cli
	validator.SetupValidator() // 验证器 Validator
	core.RunServer()
}
