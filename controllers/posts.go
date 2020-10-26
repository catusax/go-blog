package controllers

import (
	"blog/models"
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
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil && pageSize > 50 {
		pageSize = 10
	}
	posts, total, err := models.GetPublishedPostList(page, pageSize)
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
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
		_ = c.AbortWithError(404, err)
		return
	}
	post, err := models.GetPost(page)
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
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
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil && pageSize > 50 {
		pageSize = 10
	}
	archives, total, err := models.GetPostByYear(page, pageSize)
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
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
		return
	}

	posts, _, err := models.GetPublishedPostsByTag(tag, 1, 10)
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
