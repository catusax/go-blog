package models

import (
	"bytes"
	"log"
	"time"

	"github.com/spf13/viper"
)

//Page 是其他页面，如about页面等
type Page struct {
	Article
	Order int
}

//GetPagesList 获取page列表
func GetPagesList() ([]Page, int64) {
	var pages []Page
	var total int64
	db.Order("order").Find(&pages).Count(&total)
	return pages, total
}

//GetPage 根据主键获取page
func GetPage(ID int) (Page, error) {
	var page Page
	err := db.First(&page, ID).Error
	return page, err
}

//MDParse 用于把page原始Markdown解析成一个page结构体
func (page *Page) MDParse(md []byte, yaml []byte) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	if err != nil {
		log.Println(err)
	}

	page.Title = viper.GetString("title")
	if viper.GetString("date") != "" { //存在date字段说明是hexo源文件
		gettime, _ := time.Parse("2006-01-02 15:04:05", viper.GetString("date"))
		page.Update = gettime.Format("Jan 02,2006")
		page.CreatedAt = gettime
	} else {
		page.Update = time.Now().Format("Jan 02,2006")
	}
	page.Content = string(md)
	page.Yaml = string(yaml)
}

//DeletePage 根据id删除一个page
func DeletePage(id int) error {
	return db.Delete(&Page{}, id).Error
}

//Save 保存或更新一个文章到数据库
func (page *Page) Save() error {
	page.setDescription()
	page.setHTML()
	page.setDate()
	var page2 Page
	err := db.First(&page2, page.ID).Error
	//err := db.Where("title = ?", article.Title).First(&existarticle)
	//插入或更新
	if err != nil { //不存在article，直接新建
		return db.Create(&page).Error
	}
	page.CreatedAt = page2.CreatedAt //防止更新后时间错乱
	return db.Save(&page).Error
}
