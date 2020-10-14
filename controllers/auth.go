package controllers

import (
	"blog/utils"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

//用用户名和密码加密数据
var blockkey = sha256.Sum256([]byte(utils.C.User.Password))
var s = securecookie.New([]byte(utils.C.User.Username), blockkey[:])

//Login 用于处理登陆请求
func Login(c *gin.Context) {

	var user utils.User
	//log.Println(c.Request)
	c.ShouldBindJSON(&user)
	log.Println(user)
	if loginauth(user.Username, user.Password) {
		value := map[string]string{
			"key": fmt.Sprintf("%d", time.Now().Unix()),
		}
		encoded, err := s.Encode("cookies", value)
		if err != nil {
			log.Println("加密错误：", err)
		}

		cookie := &http.Cookie{
			Name:     "cookies",
			Value:    encoded,
			HttpOnly: true,
			MaxAge:   2629743,
			Secure:   true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, gin.H{
			"status":           "ok",
			"currentAuthority": "admin",
		})
	} else {
		//c.Redirect(http.StatusU, "/loginpage")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "用户名或密码错误",
		})
	}
}

// AuthMiddleWare 路由中间件，验证cookie是否正确
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		rmcookie := &http.Cookie{
			Name:     "cookies",
			Value:    "123456",
			HttpOnly: true,
			MaxAge:   0,
			Secure:   true,
			Path:     "/",
		}

		if cookie, err := c.Request.Cookie("cookies"); err == nil {

			value := make(map[string]string)
			err = s.Decode("cookies", cookie.Value, &value)
			if err != nil {
				log.Println("cookie decode error:", err) //解码失败说明cookies错误
				http.SetCookie(c.Writer, rmcookie)       //清除cookie
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  401,
					"message": "cookies已过期",
				})
				c.Abort()
				return
			}
			//log.Println("date is:", value["key"])
			cookietime, err := strconv.ParseInt(value["key"], 10, 64)
			if err != nil {
				log.Println(err)
			}

			if time.Now().Unix()-cookietime <= 2629743 { //cookies未过期
				//log.Println("Auth passed!")
				c.Next()
				return
			}
		}
		http.SetCookie(c.Writer, rmcookie) //清除cookie
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "cookies已过期",
		})
		c.Abort()
		return
	}
}

func loginauth(name string, pass string) bool {
	if name == utils.C.User.Username && pass == utils.C.User.Password {
		return true
	}

	return false
}
