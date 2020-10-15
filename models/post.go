package models

import (
	"bytes"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

// Post 包含了Tag
type Post struct {
	Article
	Tags []*Tag `gorm:"many2many:post_tags;"`
	//Tags    []Tag  `gorm:"many2many:post_tags;"`
}

//GetPublishedPostList 根据索引获取一页十个文章列表和总文章数
func GetPublishedPostList(page int, pagesize int) ([]Post, int64) {
	var posts []Post
	var total int64
	db.Preload("Tags").Select("id", "title", "update", "description").Order("created_at desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&posts, "publish = ?", true)
	db.Model(&Post{}).Where("publish = ?", true).Count(&total)
	return posts, total
}

//GetPostsList 根据参数返回对应页面和关键词的文章以及总文章数
func GetPostsList(page int, pagesize int, word string) ([]Post, int64) {
	var posts []Post
	var total int64
	if word != "" {
		db.Order("created_at desc").Limit(pagesize).Offset((page-1)*pagesize).Where("title like ?", "%"+word+"%").Find(&posts)
		db.Model(&Post{}).Where("title like ?", "%"+word+"%").Count(&total)
	} else {
		db.Preload(clause.Associations).Order("created_at desc").Limit(pagesize).Offset((page - 1) * pagesize).Find(&posts)
		db.Model(&Post{}).Count(&total)
	}
	return posts, total
}

//GetPost 根据主键获得一个post
func GetPost(ID int) (Post, error) {
	var post Post
	err := db.Preload("Tags").First(&post, ID).Error
	return post, err
}

//Archive 以年份分类保存post
type Archive struct {
	Year  int
	Posts []Post
}

//GetPostByYear 按年份分类导出post
func GetPostByYear(page int, pagesize int) ([]Archive, int64, error) {
	var posts []Post
	var archive []Archive
	var total int64

	err := db.Select("id", "title", "update", "created_at").Limit(pagesize).Offset((page-1)*pagesize).Where("publish = ?", true).Find(&posts).Error
	if err != nil {
		return nil, 0, err
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
	return archive, total, err
}

//MDParse 用于把hexo post解析成一个Post结构体
func (post *Post) MDParse(md []byte, yaml []byte) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	if err != nil {
		log.Println(err)
	}

	var tags []*Tag
	for _, tag := range viper.GetStringSlice("tags") {
		var dbtag Tag
		if db.Where("name = ?", tag).First(&dbtag).Error == nil {
			tags = append(tags, &dbtag)
		} else {
			tags = append(tags, &Tag{Name: tag})
		}
	}
	post.Tags = tags

	post.Title = viper.GetString("title")
	if viper.GetString("date") != "" { //存在date字段说明是hexo源文件
		gettime, _ := time.Parse("2006-01-02 15:04:05", viper.GetString("date"))
		post.Update = gettime.Format("Jan 02,2006")
		post.CreatedAt = gettime
	} else {
		post.Update = time.Now().Format("Jan 02,2006")
	}
	post.Content = string(md)
	post.Yaml = string(yaml)
}

//DeletePost 根据id删除一个post
func DeletePost(id int) error {
	return db.Delete(&Post{}, id).Error
}

//Save 保存或更新一个文章到数据库
func (post *Post) Save() error {
	post.setDescription()
	post.setHTML()
	post.setDate()
	var post2 Post
	err := db.First(&post2, post.ID).Error
	//err := db.Where("title = ?", article.Title).First(&existarticle)
	//插入或更新
	if err != nil { //不存在article，直接新建
		return db.Create(&post).Error
	}
	post.CreatedAt = post2.CreatedAt //防止更新后时间错乱
	return db.Save(&post).Error
}
