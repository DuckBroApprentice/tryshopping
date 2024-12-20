package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellobyte", func(context *gin.Context) {
		fullPath := "請求路徑:" + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.Write([]byte(fullPath))
	})

	engine.GET("/hellostring", func(context *gin.Context) {
		fullPath := "請求路徑:" + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.WriteString(fullPath)
	})

	engine.Run()
}
