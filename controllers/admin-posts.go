package controllers

import (
	"blog/errors"
	"blog/models"
	"blog/utils/normalize"
	"blog/utils/rss"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetPosts 返回文章列表
func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}
	word := c.Query("word")
	posts, total, err := models.GetPostsList(page, pageSize, word)
	if err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"page":  page,
		"total": total,
	})
}

//Import 导入hexo的MD文件
func Import(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		returnError(err, c)
		return
	}
	files := form.File["files"]
	if err = importFile(&files); err != nil {
		returnError(err, c)
		return
	}
	if err := rss.WriteAtom("static/atom.xml"); err != nil {
		returnError(err, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "",
	})
}

func importFile(files *[]*multipart.FileHeader) error {
	for _, fileHeader := range *files {
		file, _ := fileHeader.Open()
		fileBytes, _ := ioutil.ReadAll(file)
		normalize.LinesToLF(&fileBytes)
		_ = file.Close()
		var post models.Post
		post.Content = string(fileBytes)
		if err := post.Parse(); err != nil {
			return errors.Errorf(err, "Parse failed")
		}
		if post.Title == "" {
			return errors.New("No title found")
		}
		if err := post.NoDuplicateSave(); err != nil {
			return errors.Errorf(err, "%s", post.Title)
		}
	}
	return nil
}

//New 用于新建一个post或者更新一个post
func New(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		returnError(err, c)
		return
	} //原始文本+发布状态
	if err := post.Parse(); err != nil {
		returnError(err, c)
		return
	}
	if err := post.Save(); err != nil {
		returnError(err, c)
		return
	}
	if err := rss.WriteAtom("static/atom.xml"); err != nil {
		returnError(err, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "",
	})
}

//Delete 删除post
func Delete(c *gin.Context) {
	type d struct {
		ID int
	}
	var id d
	if err := c.ShouldBindJSON(&id); err != nil {
		returnError(err, c)
		return
	}
	if err := models.DeletePost(id.ID); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "",
	})

}

//ChangeStatus 更改文章状态
func ChangeStatus(c *gin.Context) {
	type submit struct {
		ID      int
		Publish bool
	}
	var reqJson submit
	if err := c.ShouldBindJSON(&reqJson); err != nil {
		returnError(err, c)
		return
	}
	post, err := models.GetPost(reqJson.ID)
	if err != nil {
		returnError(err, c)
		return
	}
	post.Publish = reqJson.Publish
	if err := post.Save(); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "",
	})

}
