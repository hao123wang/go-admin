package flag

import (
	"go-admin-server/global"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Auto migrate tables",
	}
	adminFlag = &cli.BoolFlag{
		Name:  "admin",
		Usage: "Crate a root account",
	}
)

func run(c *cli.Context) {
	// 不允许一条命令有多个标志，实现互斥
	if c.NumFlags() > 1 {
		global.Logger.Fatal("Only one flag can be specified")
	}
	switch {
	case c.Bool(sqlFlag.Name):
		if err := SQL(); err != nil {
			global.Logger.Fatal("Failed to automigrate", zap.Error(err))
		}
		global.Logger.Info("Successfully AutoMigrate table")
	case c.Bool(adminFlag.Name):
		if err := CreateRootAccount(); err != nil {
			global.Logger.Fatal("Failed to create root account", zap.Error(err))
		}
		global.Logger.Info("Successfully create a root account")
	default:
		global.Logger.Fatal("unknown command")
	}
}

func InitFlag() {
	if len(os.Args) > 1 {
		// 创建 app 实例
		app := cli.NewApp()
		app.Name = "go-admin-server"
		app.Flags = []cli.Flag{
			sqlFlag,
			adminFlag,
		}
		app.Action = run

		// 运行
		err := app.Run(os.Args)
		if err != nil {
			global.Logger.Error("Application execution encounted an error:", zap.Error(err))
			os.Exit(1)
		}
		os.Exit(0)
	}
}
