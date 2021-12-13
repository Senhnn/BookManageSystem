package config

import (
	"github.com/spf13/viper"
)

var (
	ServerIp   string
	ServerPort string

	JwtSigKey string
	JwtIss    string
)

// 设置默认值
func init() {
	// 设置默认地址
	viper.SetDefault("Server", map[string]string{
		"Ip":   "127.0.0.1",
		"Port": "10010",
	})

	// 设置jwt默认的自定义Payload和签名密钥
	viper.SetDefault("JWT", map[string]string{
		"JwtSigKey": "SyhanLiu19970719",
		"JwtIss":    "SyhanLiu",
	})
}

func Init() {
	viper.SetConfigName("serverCfg")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read serverCfg.yaml err!")
	}

	// 服务器地址配置
	// 服务器地址
	ServerIp = viper.GetString("Server.Ip")
	ServerPort = viper.GetString("Server.Port")

	// JWT配置
	JwtSigKey = viper.GetString("JWT.JwtSigKey")
	JwtIss = viper.GetString("JWT.JwtIss")
}
