package controllers

import (
	"blog/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Index 目录
func Index(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil && pagesize > 50 {
		pagesize = 10
	}
	posts, total := models.GetPublishedPostList(page, pagesize)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"page":  page,
		"total": total,
	})
}

//Post 用于获取单独一篇文章
func Post(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatus(404)
	}
	post, _ := models.GetPost(page)
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

//Archive 处理archive页面
func Archive(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil && pagesize > 50 {
		pagesize = 10
	}
	archives, total, err := models.GetPostByYear(page, pagesize)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"archives": archives,
		"total":    total,
	})
}

//Tag 返回TAG对应的文章列表
func Tag(c *gin.Context) {
	tag := c.Query("tag")
	if tag == "" {
		c.AbortWithStatus(404)
	}

	posts, _ := models.GetPublishedPostsByTag(tag, 1, 10)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
