package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("姓名:", person.Name)
		fmt.Println("年齡:", person.Age)
		context.Writer.Write([]byte("添加," + person.Name))

	})

	engine.Run()
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	//Age  string `form:"age"` //age為int所以報錯
	Age int `form:"age"` //變數名稱、型別要與post表單一致
}
