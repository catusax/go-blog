package models

import "blog/errors"

//Page 是其他页面，如about页面等
type Page struct {
	Article
	MenuName string `gorm:"type:varchar(15)"`
	Comment  bool
	Enable   bool
}

//GetPagesList 获取page列表
func GetPagesList() ([]Page, error) {
	var pages []Page
	err := db.Order("id").Find(&pages).Error
	return pages, errors.Errorf(err, "Database query failed")
}

//GetPage 根据主键获取page
func GetPage(ID int) (Page, error) {
	var page Page
	err := db.First(&page, ID).Error
	return page, errors.Errorf(err, "Database query failed")
}

////DeletePage 根据id删除一个page
//func DeletePage(id int) error {
//	return db.Delete(&Page{}, id).Error
//}

//Save 保存或更新一个文章到数据库
func (page *Page) Save() error {
	page.setHTML()
	page.setDate()
	var page2 Page
	err := db.First(&page2, page.ID).Error
	//插入或更新
	if err != nil { //不存在article，直接新建
		return errors.Errorf(db.Create(&page).Error, "Database insert failed")
	}
	page.CreatedAt = page2.CreatedAt //防止更新后时间错乱
	return errors.Errorf(db.Save(&page).Error, "Database update failed")
}
