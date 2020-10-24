package controllers

import (
	"blog/utils"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//UploadFavicon
func UploadFavicon(c *gin.Context) {
	file, err := c.FormFile("favicon")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}
	err = c.SaveUploadedFile(file, "./static/favicon.ico")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

//UploadAvatar
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}
	err = c.SaveUploadedFile(file, "./static/avatar.png")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
//ChangeConfig
func ChangeConfig(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	err := utils.WriteConf(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

//GetConfig
func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"SiteName": utils.C.SiteName,
		"Port":     utils.C.Port,
		"DB":       viper.GetStringMap("DB"),
		"Disqus":   viper.GetStringMap("disqus"),
	})
}

//Info
func Info(c *gin.Context) {
	time.Sleep(1 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"SiteName": utils.C.SiteName,
		"Disqus":   viper.GetStringMap("disqus"),
	})
}
