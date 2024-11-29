package main

import (
	"github.com/DuckBroApprentice/Shopping/database"
	"github.com/DuckBroApprentice/Shopping/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	routers.AddUserRouter(v1)

	go func() {
		database.ConnectToMySQL()
	}()

	router.Run(":8080")
}
