package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 路由跟入口尚未拆開
func main() {
	engine := gin.Default()

	//localhost:8080/hello?name=davie
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		//輸出
		context.Writer.Write([]byte("Hello" + name))

	})
	//post
	//8080/login
	//這是簡易的登入，要再加上帳密比對
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath()) //打印請求接口

		//解析前端的輸入 post表單提交的字段
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)

		context.Writer.Write([]byte(username + "登錄成功"))
	})

	engine.Run()
}
