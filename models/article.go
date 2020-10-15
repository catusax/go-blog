package models

import (
	"blog/utils"
	"time"

	"github.com/russross/blackfriday/v2"
	"gorm.io/gorm"
)

//Article 是基本的文章结构
type Article struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null"`
	Content     string `gorm:"type:text"`
	HTML        string `gorm:"type:text"`
	Update      string `gorm:"type:varchar(12);default:2020-09-28"`
	Description string `gorm:"type:varchar(400)"`
	Publish     bool   `gorm:"default:false"`
	Yaml        string `gorm:"type:text"`
}

//SetDescription 根据文章Content生成Description
func (article *Article) setDescription() {
	descMD := utils.GetDescription([]byte(article.Content))
	if len(descMD) >= 5 {
		article.Description = string(blackfriday.Run(descMD))
	}
}

//将MD转换为HTML
func (article *Article) setHTML() {
	article.HTML = string(blackfriday.Run([]byte(article.Content)))
}

func (article *Article) setDate() {
	article.UpdatedAt = time.Now()
	article.Update = time.Now().Format("Jan 02,2006")
}
