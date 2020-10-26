package models

import (
	"blog/utils"
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", utils.C.DB.Host, utils.C.DB.User, utils.C.DB.Password, utils.C.DB.Name, utils.C.DB.Port)
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
