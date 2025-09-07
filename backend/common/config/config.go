// 读取 Yaml 文件

package config

import "github.com/spf13/viper"

type AppConfig struct {
	Server      `mapstructure:"server"`
	Mysql       `mapstructure:"mysql"`
	Redis       `mapstructure:"redis"`
	Logger      `mapstructure:"logger"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type Mysql struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DB           string `mapstructure:"db"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	LogLevel     string `mapstructure:"log_level"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Logger struct {
	Filename       string `mapstructure:"filename"`
	Level          string `mapstructure:"level"`
	MaxSize        int    `mapstructure:"max_size"`
	MaxAge         int    `mapstructure:"max_age"`
	MaxBackups     int    `mapstructure:"max_backups"`
	IsConsolePrint bool   `mapstructure:"is_console_print"`
}

func Init() *AppConfig {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
