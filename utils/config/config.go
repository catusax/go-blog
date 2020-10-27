package config

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

//Config 配置文件
type Config struct {
	Port     int
	SiteName string
	DB       DB
	User     User
}

//DB 数据库配置
type DB struct {
	Host     string
	Port     int
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

//WriteConf 将更改后的配置写入配置文件
func WriteConf(data []byte) error {
	var config map[string]interface{}
	//json.NewDecoder(map1).Decode(&b)
	err := json.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	for k, v := range config {
		viper.Set(k, v)
	}
	viper.Unmarshal(&C)
	viper.WriteConfig()
	return nil
}
