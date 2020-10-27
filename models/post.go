package models

import (
	"blog/errors"
	"blog/utils/md"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

// Post 包含了Tag
type Post struct {
	Article
	Description string `gorm:"type:text"`
	Publish     bool   `gorm:"default:false"`
	Tags        []*Tag `gorm:"many2many:post_tags;"`
}

//GetPublishedPostList 根据索引获取一页十个文章列表和总文章数
func GetPublishedPostList(page int, pageSize int) ([]Post, int64, error) {
	var posts []Post
	var total int64
	err := db.Preload("Tags").Select("id", "title", "update", "description").Order("created_at desc").Limit(pageSize).Offset((page-1)*pageSize).Find(&posts, "publish = ?", true).Error
	db.Model(&Post{}).Where("publish = ?", true).Count(&total)
	return posts, total, errors.Errorf(err, "Database query failed")
}

// GetAllPublishedPost 根据索引获取一页十个文章列表和总文章数
func GetAllPublishedPost() ([]Post, error) {
	var posts []Post
	//var total int64
	err := db.Preload("Tags").Select("id", "title", "update", "description").Order("created_at desc").Find(&posts, "publish = ?", true).Error
	//db.Model(&Post{}).Where("publish = ?", true).Count(&total)
	return posts, errors.Errorf(err, "Database query failed")
}

//GetPostsList 根据参数返回对应页面和关键词的文章以及总文章数
func GetPostsList(page int, pageSize int, word string) ([]Post, int64, error) {
	var posts []Post
	var total int64
	var err error
	if word != "" {
		err = db.Order("created_at desc").Limit(pageSize).Offset((page-1)*pageSize).Where("title like ?", "%"+word+"%").Find(&posts).Error
		db.Model(&Post{}).Where("title like ?", "%"+word+"%").Count(&total)
	} else {
		err = db.Preload(clause.Associations).Order("created_at desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts).Error
		db.Model(&Post{}).Count(&total)
	}
	return posts, total, errors.Errorf(err, "Database failed")
}

//GetPost 根据主键获得一个post
func GetPost(ID int) (Post, error) {
	var post Post
	err := db.Preload("Tags").First(&post, ID).Error
	return post, errors.Errorf(err, "Database query failed")
}

//Archive 以年份分类保存post
type Archive struct {
	Year  int
	Posts []Post
}

//GetPostByYear 按年份分类导出post
//goland:noinspection GoNilness
func GetPostByYear(page int, pageSize int) ([]Archive, int64, error) {
	var posts []Post
	var archive []Archive
	var total int64

	err := db.Select("id", "title", "update", "created_at").Limit(pageSize).Offset((page-1)*pageSize).Where("publish = ?", true).Find(&posts).Error
	if err != nil {
		return nil, 0, errors.Errorf(err, "Database query failed")
	}
	db.Model(&Post{}).Count(&total)
	year := 0 //上一年年份
	num := -1 //数组位置
	for _, post := range posts {
		//获取时间
		date := post.CreatedAt
		//年份改变就追加一个archive
		if date.Year() != year {
			year = date.Year()
			num++
			archive = append(archive, Archive{Year: year, Posts: nil})
			archive[num].Posts = append(archive[num].Posts, post)
		} else { //年份不变就追加一个post
			archive[num].Posts = append(archive[num].Posts, post)
		}
	}
	return archive, total, nil
}

//SetDescription 根据文章Content生成Description
func (post *Post) setDescription() {
	descMD := md.GetDescription([]byte(post.Content))
	if len(descMD) >= 5 {
		post.Description = string(markdown.ToHTML(descMD, nil, nil))
	}
}

//MDParse 用于把hexo post解析成一个Post结构体
func (post *Post) setTags() error {
	var viperTags = viper.New()
	viperTags.SetConfigType("yaml")
	err := viperTags.ReadConfig(strings.NewReader(post.Yaml))
	if err != nil {
		return errors.Errorf(err, "viper ReadConfig failed")
	}

	var tags []*Tag
	for _, tag := range viperTags.GetStringSlice("tags") {
		var dbTag Tag
		if db.Where("name = ?", tag).First(&dbTag).Error == nil {
			tags = append(tags, &dbTag)
		} else {
			tags = append(tags, &Tag{Name: tag})
		}
	}
	post.Tags = tags
	return nil
}

//DeletePost 根据id删除一个post
func DeletePost(id int) error {
	return errors.Errorf(db.Select("posts", "tags").Delete(&Post{}, id).Error, "Database delete failed")
}

//Save 保存或更新一个文章到数据库
func (post *Post) Save() error {
	_ = post.setTags()
	post.setDescription()
	post.setHTML()
	post.setDate()
	var post2 Post
	err := db.First(&post2, post.ID).Error
	//插入或更新
	if err != nil { //不存在article，直接新建
		return errors.Errorf(db.Create(&post).Error, "Database insert failed")
	}
	post.CreatedAt = post2.CreatedAt //防止更新后时间错乱
	return errors.Errorf(db.Save(&post).Error, "Database update failed")
}

// NoDuplicateSave 保存文章时候保证标题不重复,用于从hexo导入文章
func (post *Post) NoDuplicateSave() error {
	_ = post.setTags()
	post.setDescription()
	post.setHTML()
	post.Publish = true
	var p Post
	if err := db.Where("title = ?", post.Title).First(&p).Error; err != nil {
		if err := db.Create(&post).Error; err != nil {
			return errors.Errorf(err, "Database insert failed")
		}
		return nil
	}
	return errors.New("duplicate title")
}
