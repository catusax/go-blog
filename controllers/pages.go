package controllers

import (
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PagesList 返回Page列表
func PagesList(c *gin.Context) {
	pages, err := models.GetPagesList()
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})

}

//Page 根据ID获得某个Page
func Page(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
	page, err := models.GetPage(ID)
	if err != nil {
		_ = c.AbortWithError(404, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page": page,
	})
}

//NewPage 新建或更新一个page
func NewPage(c *gin.Context) {
	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		returnError(err, c)
		return
	}
	if err := page.Save(); err != nil {
		returnError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "",
	})
}
