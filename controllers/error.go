package controllers

import "github.com/gin-gonic/gin"

func returnError(err error, c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "error",
		"msg":    err.Error(),
	})
	_ = c.Error(err)
}
