package models

import "gorm.io/gorm"

//Tag 存储tags
type Tag struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(10);not null"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}

//GetPostsByTag 获得tag对应的文章列表
func GetPostsByTag(name string, page int, pagesize int) ([]Post, int64) {
	var tag Tag
	var posts []Post
	db.Where("name = ?", name).First(&tag)
	db.Model(&tag).Limit(pagesize).Offset((page - 1) * pagesize).Association("Posts").Find(&posts)
	total := db.Model(&tag).Association("Posts").Count()
	return posts, total
}

//GetPublishedPostsByTag 获得tag对应的文章列表
func GetPublishedPostsByTag(name string, page int, pagesize int) ([]Post, int64) {
	var tag Tag
	var posts []Post
	//db.Preload("Posts", "publish = ?", true).First(&tag, "name = ?", name)
	db.Where("name = ?", name).First(&tag)
	db.Model(&tag).Limit(pagesize).Offset((page-1)*pagesize).Where("publish = ?", true).Association("Posts").Find(&posts)
	total := db.Model(&tag).Where("publish = ?", true).Association("Posts").Count()
	return posts, total
}

//GetTagList 获取tag列表
func GetTagList() []Tag {
	var tags []Tag
	db.Find(&tags)
	return tags
}

//DeleteTag 根据tag名称删除tag
func DeleteTag(name string) error {
	return db.Where("name = ?", name).Delete(&Tag{}).Error
}
