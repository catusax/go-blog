package models

import (
	"blog/utils/config"
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.C.DB.Host, config.C.DB.User, config.C.DB.Password, config.C.DB.Name, config.C.DB.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}) //初始化log
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	err = db.AutoMigrate(&Post{}, &Page{}, &Visitor{})
	if err != nil {
		log.Fatal("Failed to migrate data: ", err)
	}
}
