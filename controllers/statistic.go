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

func RecentPost(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetRecentPost())
}

func RecentVisit(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetRecentVisit())
}

func MostReaded(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetMostReadedPost())
}

func Browsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetBrowser())
}

func OS(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetOS())
}
