package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		var register Register
		fmt.Println(context.FullPath())

		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)

	})

	engine.Run()
}

type Register struct {
	UserName string
	Phone    string
	Password string
}
