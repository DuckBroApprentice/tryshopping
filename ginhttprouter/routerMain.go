package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//localhost:8080/hello?name=davie
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello," + name))
	})

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		//GetPostForm就能比對資料
		//與SQL連結比對資料?
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}
		password, exist := context.GetPostForm("password")
		if exist {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("hello ," + username))
	})

	//id是個變值，無法確定，要起個變量名 ex. :id
	//此做法: 8080/user/123  -> delete 用戶ID 123
	//可以改成post表單來處理，並加上比對是否存在此請求ID
	engine.DELETE("user/:id", func(context *gin.Context) {
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("delete 用戶ID" + userID))
	})

	engine.Run()
}
