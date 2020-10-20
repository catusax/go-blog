package controllers

import (
	"blog/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PagesList 返回Page列表
func PagesList(c *gin.Context) {
	pages := models.GetPagesList()
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})

}

//Page 根据ID获得某个Page
func Page(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatus(404)
	}
	page, err := models.GetPage(ID)
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(http.StatusOK, gin.H{
		"page": page,
	})
}

//NewPage 新建或更新一个page
func NewPage(c *gin.Context) {
	var page models.Page
	c.ShouldBindJSON(&page)
	log.Println("解析结果：", page)
	err := page.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "",
		})
	}
}
