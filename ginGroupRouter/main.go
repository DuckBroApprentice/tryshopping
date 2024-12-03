package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	routerGroup := engine.Group("/user") //第一級/user
	routerGroup.POST("/register", registerHandle)
	routerGroup.POST("/login", loginHandle)
	routerGroup.DELETE("/:id", deleteHandle)

	engine.Run()
}

func registerHandle(c *gin.Context) {
	fullPath := "用戶註冊:" + c.FullPath()
	fmt.Println(fullPath)
	c.Writer.WriteString(fullPath)
}

func loginHandle(c *gin.Context) {
	fullPath := "用戶登錄" + c.FullPath()
	fmt.Println(fullPath)
	c.Writer.WriteString(fullPath)
}

func deleteHandle(c *gin.Context) {
	fullPath := "用戶刪除" + c.FullPath()
	userID := c.Param("id")
	fmt.Println(fullPath + " " + userID)
	c.Writer.WriteString(fullPath + " " + userID)
}
