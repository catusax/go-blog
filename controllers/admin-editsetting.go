package controllers

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
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
	c.JSON(http.StatusOK, viper.AllSettings())
}
