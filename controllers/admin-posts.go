package controllers

import (
	"blog/models"
	"blog/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// posts.GET("/getlist", ctrs.GetPosts)
// posts.POST("/import")
// posts.POST("/new")
// posts.PUT("/update")
// posts.DELETE("/delete")

//GetPosts 返回文章列表
func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pagesize = 10
	}
	word := c.Query("word")
	posts, total := models.GetPostsList(page, pagesize, word)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"page":  page,
		"total": total,
	})
}

//Import 导入hexo的MD文件
func Import(c *gin.Context) {
	file, _, _ := c.Request.FormFile("file")
	fileBytes, err := ioutil.ReadAll(file) //读取内容
	if err != nil {
		log.Println(err)
	}
	yaml, md := utils.MDCut(fileBytes)
	var post models.Post
	post.MDParse(md, yaml)
	err = post.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "",
		})
	}
}

//New 用于新建一个post或者更新一个post
func New(c *gin.Context) {
	var post models.Post
	c.ShouldBindJSON(&post) //原始文本+发布状态
	yaml, md := utils.MDCut([]byte(post.Content))
	post.MDParse(md, yaml)
	err := post.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "",
		})
	}
}

//Delete 删除post
func Delete(c *gin.Context) {
	type delete struct {
		ID int
	}
	var id delete
	c.ShouldBindJSON(&id)
	err := models.DeletePost(id.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "",
		})
	}
}

//ChangeStatus 更改文章状态
func ChangeStatus(c *gin.Context) {
	type submit struct {
		ID      int
		Publish bool
	}
	var reqjson submit
	c.ShouldBindJSON(&reqjson)
	post, err := models.GetPost(reqjson.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		post.Publish = reqjson.Publish
		post.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "",
		})
	}
}
