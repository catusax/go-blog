package controllers

import (
	"blog/utils/config"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// UploadFavicon 上传favicon
func UploadFavicon(c *gin.Context) {
	file, err := c.FormFile("favicon")
	if err != nil {
		returnError(err, c)
		return
	}
	if err := c.SaveUploadedFile(file, "./www/favicon.ico"); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		returnError(err, c)
		return
	}
	if err := c.SaveUploadedFile(file, "./www/avatar.png"); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// ChangeConfig 修改配置文件
func ChangeConfig(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := config.WriteConf(data); err != nil {
		returnError(err, c)
		return
	}
	config.ReadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetConfig 获取配置文件
func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}
