package main

import (
	"io"
	"os"

	"blog/routers"
	"blog/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 设置日志文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultWriter = io.Writer(os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	//cors
	corsconf := cors.DefaultConfig()
	corsconf.AllowCredentials = true
	corsconf.AllowOrigins = []string{"http://localhost:8000", "http://localhost:8001", "http://127.0.0.1", "http://blog.coolrc.me"}

	r.Use(cors.New(corsconf))
	// 加载路由
	routers.LoadRouters(r)
	r.Run(":" + utils.C.Port)
}
