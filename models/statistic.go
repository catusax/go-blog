package models

import (
	"github.com/ua-parser/uap-go/uaparser"
	"gorm.io/gorm"
)

var Parser *uaparser.Parser

func init() {
	Parser = uaparser.NewFromSaved()
}

type Visitor struct {
	gorm.Model
	ID      uint `gorm:"primarykey"`
	PostID  int
	IP      string `gorm:"type:varchar(45)"`
	OS      string `gorm:"type:varchar(20)"`
	Browser string `gorm:"type:varchar(20)"`
}

func (v *Visitor) Create() error {
	return db.Create(&v).Error
}
func (v *Visitor) SetUA(ua string) {
	client := Parser.Parse(ua)
	v.OS = client.Os.Family
	v.Browser = client.UserAgent.Family
}

// GetTotalStatistic返回总PV和文章数
func GetTotalStatistic() (totalvisitor int64, totalpost int64) {
	db.Table("visitors").Count(&totalvisitor)
	db.Table("posts").Count(&totalpost)
	return
}

type RecentVisitor struct {
	Date  string
	Count int
}

//GetRecentVisit 获取最近一个月的访客数
func GetRecentVisit() (visitor []RecentVisitor) {
	db.Raw("select date(updated_at),count(date(updated_at)) from visitors where date_part('day',now()-updated_at)<=30 group by date order by date").Scan(&visitor)
	return
}

type RecentPost struct {
	ID     uint
	Update string
	Title  string
}

//GetRecentPost 最近更新的十篇文章
func GetRecentPost() (posts []RecentPost) {
	db.Table("posts").Select("id", "title", "update").Order("updated_at desc").Limit(10).Scan(&posts)
	return
}

type PostRead struct {
	Title  string
	Update string
	PostID uint
	Count  int
}

//GetMostReadPost 最多阅读量的十篇文章
func GetMostReadPost() (posts []PostRead) {
	db.Raw("select v.post_id ,p.title,p.update,count(v.post_id) from visitors v inner join posts p on p.id=v.post_id group by v.post_id ,p.title,p.update order by count desc limit 10").Scan(&posts)
	return
}

type BrowserTable struct {
	Browser string
	Count   int
}

//GetBrowsers 不同浏览器的访问量
func GetBrowser() (browsers []BrowserTable) {
	db.Table("visitors").Select("browser", "count(browser)").Group("browser").Order("count desc").Scan(&browsers)
	return
}

type OSTable struct {
	Os    string
	Count int
}

//GetOS 不同设备访问量
func GetOS() (os []OSTable) {
	db.Table("visitors").Select("os", "count(os)").Group("os").Order("count desc").Scan(&os)
	return
}
