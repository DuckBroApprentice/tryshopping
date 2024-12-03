package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engines := gin.Default()
	engines.Use(RequestInfos())
	/*
		//engines.Use(RequestInfos()) /hello就不會執行中間件

		//9001/query
		engines.GET("/query", RequestInfos(), func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"code": 1,
				"msg":  c.FullPath(),
			})
		})
	*/
	engines.GET("/query", func(c *gin.Context) {
		fmt.Println("中間件的使用方法")
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  c.FullPath(),
		})
	})
	engines.GET("/hello", func(c *gin.Context) {
		//todo
	})

	engines.Run(":9001")
}

// 打印請求信息的中間件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("請求Path:", path)
		fmt.Println("請求method:", method)
		fmt.Println("狀態碼:", context.Writer.Status())

		//fmt.Println("狀態碼:", context.Writer.Status())  無法獲得http.Status
		context.Next() //以上的代碼及以下的代碼分兩部分

		fmt.Println("狀態碼:", context.Writer.Status())
	}
}
