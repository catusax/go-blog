package models

//Page 是其他页面，如about页面等
type Page struct {
	Article
	MenuName string `gorm:"type:varchar(15)"`
	Comment  bool
	Enable   bool
}

//GetPagesList 获取page列表
func GetPagesList() []Page {
	var pages []Page
	db.Order("id").Find(&pages)
	return pages
}

//GetPage 根据主键获取page
func GetPage(ID int) (Page, error) {
	var page Page
	err := db.First(&page, ID).Error
	return page, err
}

//DeletePage 根据id删除一个page
func DeletePage(id int) error {
	return db.Delete(&Page{}, id).Error
}

//Save 保存或更新一个文章到数据库
func (page *Page) Save() error {
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
