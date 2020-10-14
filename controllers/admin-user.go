package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CurrentUser 返回当前用户信息
func CurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "管理员",
		"avatar":    "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"userid":    "00000001",
		"email":     "antdesign@alipay.com",
		"signature": "海纳百川，有容乃大",
	})
}
