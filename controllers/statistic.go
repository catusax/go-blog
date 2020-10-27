package controllers

import (
	"blog/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Visitor 记录访客信息
func Visitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			var visit models.Visitor
			visit.PostID, _ = strconv.Atoi(c.Query("page"))
			visit.IP = cCp.Request.Header.Get("X-Forwarded-For")
			if visit.IP == "" {
				visit.IP = c.ClientIP()
			}
			visit.SetUA(cCp.Request.Header.Get("User-Agent"))
			err := visit.Create()
			if err != nil {
				log.Println(err)
			}
		}()
		c.Next()
	}
}

// TotalStatistic 总访客数和文章数统计
func TotalStatistic(c *gin.Context) {
	visitors, posts := models.GetTotalStatistic()
	c.JSON(http.StatusOK, gin.H{
		"Visitors": visitors,
		"Posts":    posts,
	})
}

// RecentPost 最近更新的文章
func RecentPost(c *gin.Context) {
	data, err := models.GetRecentPost()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

// RecentVisit 最近一个月访客
func RecentVisit(c *gin.Context) {
	data, err := models.GetRecentVisit()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

// MostRead 最多阅读的十篇文章
func MostRead(c *gin.Context) {
	data, err := models.GetMostReadPost()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

// Browsers 浏览器分析
func Browsers(c *gin.Context) {
	data, err := models.GetBrowser()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

// OS 操作系统分析
func OS(c *gin.Context) {
	data, err := models.GetOS()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}
