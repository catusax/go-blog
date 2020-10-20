package routers

import (
	ctrs "blog/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 设置静态文件夹,仅用于测试，部署时请使用nginx or caddy
	router.Static("/font", "./static/font")
	router.StaticFile("/favicon.png", "./static/favicon.png")
	router.StaticFile("/", "./web/front/dist/index.html")
	router.StaticFile("/umi.css", "./web/front/dist/umi.css")
	router.StaticFile("/umi.js", "./web/front/dist/umi.js")
	router.Static("/admin", "./web/admin/dist")

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		posts.Use(ctrs.AuthMiddleWare()) //API使用authmiddleware验证身份
		{
			posts.GET("/getlist", ctrs.GetPosts)
			posts.POST("/import", ctrs.Import)
			posts.POST("/new", ctrs.New)
			posts.PUT("/update", ctrs.New)
			posts.DELETE("/delete", ctrs.Delete)
			posts.PUT("/changestatus", ctrs.ChangeStatus)
		}

		pages := api.Group("/pages")
		{
			pages.GET("/getlist", ctrs.PagesList)
			pages.GET("/page", ctrs.Page)
			pages.POST("/new", ctrs.NewPage)
		}

		api.GET("/currentUser", ctrs.AuthMiddleWare(), ctrs.CurrentUser)

		public := api.Group("/public")
		{
			public.GET("/index", ctrs.Index)
			public.GET("/archives", ctrs.Archive)
			public.GET("/post", ctrs.Post)
			public.POST("/login", ctrs.Login)
			public.GET("/tag", ctrs.Tag)
		}
	}
}
