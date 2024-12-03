package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//設置html目錄(根目錄)
	engine.LoadHTMLGlob("./html/*")
	//設置靜態資源目錄
	//1 請求路徑   2 root路徑(本地工程)
	engine.Static("/img", "./img")
	engine.GET("/hellohtml", func(c *gin.Context) {
		fullPath := c.FullPath()
		fmt.Println(fullPath)
		//index.html為路徑
		//第三個參數用於渲染模版
		c.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath, //key為.html變量 val為.go變數
			"title":    "gin教程",
		})
	})

	engine.Run()
}
