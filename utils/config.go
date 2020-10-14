package utils

import (
	"log"

	"github.com/spf13/viper"
)

//Config 配置文件
type Config struct {
	Port string
	DB   DB
	User User
}

//DB 数据库配置
type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

//User 用户
type User struct {
	Username string
	Password string
}

//C 存储全局配置
var C Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatal("解析配置文件失败：", err)
	}

}
