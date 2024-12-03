package main

import (
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/controller"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	registerRouter(app)

	//app.json中app_host,app_port
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 路由設置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
}
