package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//8080/hello?name=davie&classes=軟件工程
	engine.GET("hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		//綁定方式shouldbindquery
		//映射student Student
		var student Student //滿足ShouldBindQuery參數要求
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		context.Writer.Write([]byte("hello," + student.Name))
	})
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
