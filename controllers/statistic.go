package controllers

import (
	"blog/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func TotalStatistic(c *gin.Context) {
	visitors, posts := models.GetTotalStatistic()
	c.JSON(http.StatusOK, gin.H{
		"Visitors": visitors,
		"Posts":    posts,
	})
}

func RecentPost(c *gin.Context) {
	data, err := models.GetRecentPost()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

func RecentVisit(c *gin.Context) {
	data, err := models.GetRecentVisit()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

func MostRead(c *gin.Context) {
	data, err := models.GetMostReadPost()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

func Browsers(c *gin.Context) {
	data, err := models.GetBrowser()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}

func OS(c *gin.Context) {
	data, err := models.GetOS()
	if err != nil {
		_ = c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, data)
}
