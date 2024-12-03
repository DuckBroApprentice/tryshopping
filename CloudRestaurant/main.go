package main

import (
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	//app.json中app_host,app_port
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
