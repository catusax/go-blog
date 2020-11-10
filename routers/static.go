package routers

import (
	"github.com/gin-gonic/gin"
	"os"
)

func loadStatic(router *gin.Engine) {
	// 把不认识的路由全指向到index.html文件
	router.NoRoute(func(c *gin.Context) {
		c.File("www/blog/index.html")
	})
	router.StaticFile("/", "www/blog/index.html")
	router.StaticFile("/umi.css", "www/blog/umi.css")
	router.StaticFile("/umi.js", "www/blog/umi.js")
	router.StaticFile("/favicon.ico", "www/favicon.ico")
	router.StaticFile("/avatar.png", "www/avatar.png")

	router.Static("/static", "www/static")

	// /admin的rewrite规则
	router.StaticFile("/admin", "www/admin/index.html")
	TryFiles(router, "/admin/*file", "www/admin", "www/admin/index.html")
}

// TryFiles 类似nginx的try_files，用于服务SPA
func TryFiles(group *gin.Engine, relativePath, filePath, fallbackFile string) {
	handler := func(c *gin.Context) {
		file := filePath + c.Param("file")
		if isExist(file) {
			c.File(file)
		} else {
			c.File(fallbackFile)
		}
	}
	group.GET(relativePath, handler)
	group.HEAD(relativePath, handler)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
