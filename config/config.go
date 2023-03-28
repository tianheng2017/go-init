package config

import (
	"flag"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

var Config = loadConfig(".")

// envConfig 环境配置
type envConfig struct {
	GinMode    string `mapstructure:"GIN_MODE"`    // gin运行模式
	PublicUrl  string `mapstructure:"PUBLIC_URL"`  // 对外发布的Url
	ServerPort int    `mapstructure:"SERVER_PORT"` // 服务运行端口
	Version    string // 版本
	Secret     string // 系统加密字符
}

// loadConfig 加载配置
func loadConfig(path string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file path.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(path)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()
	config := envConfig{
		GinMode: "debug",
		// 服务运行端口
		ServerPort: 8000,
		// 全局配置
		// 版本
		Version: "v1.0.0",
		// 系统加密字符
		Secret: "UVTIyzCy",
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	// PublicUrl未设置设置默认值
	if config.PublicUrl == "" {
		config.PublicUrl = "http://127.0.0.1:" + strconv.Itoa(config.ServerPort)
	}
	return config
}
