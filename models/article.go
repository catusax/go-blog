package models

import (
	"blog/errors"
	"blog/utils/md"
	"bytes"
	"github.com/spf13/viper"
	"log"
	"time"

	"github.com/gomarkdown/markdown"
	"gorm.io/gorm"
)

//Article 是基本的文章结构
type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string `gorm:"type:text"`
	HTML    string `gorm:"type:text"`
	Update  string `gorm:"type:varchar(12);default:2020-09-28"`
	Yaml    string `gorm:"type:text"`
}

func (article *Article) Parse() error {
	if err := article.hexoParse(md.Cut([]byte(article.Content))); err != nil {
		return errors.Errorf(err, "Hexo Parse failed")
	}
	return nil
}

//将MD转换为HTML
func (article *Article) setHTML() {
	article.HTML = string(markdown.ToHTML([]byte(article.Content), nil, nil))
}

func (article *Article) setDate() {
	article.UpdatedAt = time.Now()
	article.Update = time.Now().Format("Jan 02,2006")
}

//MDParse 用于把hexo post解析成一个Post结构体
func (article *Article) hexoParse(yaml []byte, md []byte) error {
	article.Content = string(md)
	article.Yaml = string(yaml)
	var viperMd = viper.New()
	viperMd.SetConfigType("yaml")
	err := viperMd.ReadConfig(bytes.NewBuffer(yaml))
	if err != nil {
		return errors.Errorf(err, "viper ReadConfig failed")
	}

	article.Title = viperMd.GetString("title")
	if date := viperMd.GetString("date"); date != "" { //存在date字段说明是hexo源文件
		log.Println("Found Hexo post!date:", date)
		getTime, _ := time.Parse("2006-01-02 15:04:05", date)
		article.Update = getTime.Format("Jan 02,2006")
		article.CreatedAt = getTime
		article.UpdatedAt = getTime
	} else {
		article.Update = time.Now().Format("Jan 02,2006")
	}
	return nil
}
