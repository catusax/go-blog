package controllers

import (
	"blog/utils/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// UploadFavicon
func UploadFavicon(c *gin.Context) {
	file, err := c.FormFile("favicon")
	if err != nil {
		returnError(err, c)
		return
	}
	if err := c.SaveUploadedFile(file, "./static/favicon.ico"); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UploadAvatar
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		returnError(err, c)
		return
	}
	if err := c.SaveUploadedFile(file, "./static/avatar.png"); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// ChangeConfig
func ChangeConfig(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := config.WriteConf(data); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetConfig
func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}
