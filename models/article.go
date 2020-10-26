package models

import (
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

//将MD转换为HTML
func (article *Article) setHTML() {
	article.HTML = string(markdown.ToHTML([]byte(article.Content), nil, nil))
}

func (article *Article) setDate() {
	article.UpdatedAt = time.Now()
	article.Update = time.Now().Format("Jan 02,2006")
}
