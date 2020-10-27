package controllers

import (
	"blog/utils/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// Info
func Info(c *gin.Context) {
	//time.Sleep(1 * time.Second) //测试
	c.JSON(http.StatusOK, gin.H{
		"SiteName": config.C.SiteName,
		"Disqus":   viper.GetStringMap("disqus"),
		"Github":   viper.GetString("github"),
	})
}
