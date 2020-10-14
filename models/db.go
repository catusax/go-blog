package models

import (
	"blog/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", utils.C.DB.Host, utils.C.DB.User, utils.C.DB.Password, utils.C.DB.Name, utils.C.DB.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{ /*Logger: logger.Default.LogMode(logger.Info)*/ }) //初始化log
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	err = db.AutoMigrate(&Post{})
	if err != nil {
		log.Fatal(err)
	}
}

//TODO 简介字段
func creat(title string, content string) error {
	result := db.Create(map[string]interface{}{
		"Title": title, "Content": content,
	})
	return result.Error
}

//读取整页十个 读取文章内容
func read() {

}
