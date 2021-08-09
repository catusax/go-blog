package main

import (
	"blog/utils/config"
	"io"
	"log"
	"os"
	"strconv"

	"blog/routers"
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
	// cors
	corsConf := cors.DefaultConfig()
	corsConf.AllowCredentials = true
	corsConf.AllowOrigins = []string{"http://localhost:8000", "http://localhost:8001", "http://127.0.0.1", "http://blog.coolrc.me"}
	corsConf.AllowHeaders = []string{"x-requested-with", "content-type"}

	r.Use(cors.New(corsConf))
	// 加载路由
	routers.LoadRouters(r)
	log.Println("Starting server on :",config.C.Port)
	err := r.Run(":" + strconv.Itoa(config.C.Port))
	if err != nil {
		log.Panic("err:",err)
	}
}
