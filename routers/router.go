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

//nolint:funlen
func loadRouters(router *gin.Engine) {

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

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

		setting := api.Group("/settings")
		setting.Use(ctrs.AuthMiddleWare())
		{
			setting.POST("/avatar", ctrs.UploadAvatar)
			setting.POST("/favicon", ctrs.UploadFavicon)
			setting.POST("/changeconfig", ctrs.ChangeConfig)
			setting.GET("/getconfig", ctrs.GetConfig)
		}

		statistic := api.Group("/statistic")
		statistic.Use(ctrs.AuthMiddleWare())
		{
			statistic.GET("/total", ctrs.TotalStatistic)
			statistic.GET("/recentpost", ctrs.RecentPost)
			statistic.GET("/recentvisit", ctrs.RecentVisit)
			statistic.GET("/mostread", ctrs.MostRead)
			statistic.GET("/browser", ctrs.Browsers)
			statistic.GET("/os", ctrs.OS)
		}

		api.GET("/currentUser", ctrs.AuthMiddleWare(), ctrs.CurrentUser)

		public := api.Group("/public")
		{
			public.GET("/index", ctrs.Index)
			public.GET("/archives", ctrs.Archive)
			public.GET("/post", ctrs.Visitor(), ctrs.Post)
			public.POST("/login", ctrs.Login)
			public.GET("/tag", ctrs.Tag)
			public.GET("/info", ctrs.Info)
		}
	}
}
