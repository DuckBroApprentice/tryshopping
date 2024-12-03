package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellojson", func(c *gin.Context) {
		fullPath := "請求路徑:" + c.FullPath()
		fmt.Println(fullPath)
		c.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    fullPath,
		})
	})

	engine.GET("/jsonstruct", func(c *gin.Context) {
		fullPath := "請求路徑:" + c.FullPath()
		fmt.Println(fullPath)

		resp := Response{Code: 1, Message: "OK", Data: fullPath}

		c.JSON(200, &resp)
	})

	engine.Run()
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
