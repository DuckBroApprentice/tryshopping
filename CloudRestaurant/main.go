package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/controller"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//初始化redis配置
	tool.InitRedisStore()

	app := gin.Default()

	//設置全局跨域訪問
	app.Use(Cors())

	registerRouter(app)

	//app.json中app_host,app_port
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 路由設置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}

// 跨域訪問:cross origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin") //header內的Origin
		//解析head其他內容
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*")                                   //"*" :允許訪問所有域
			context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE") //允許訪問的方法
			context.Header("Access-Control-Allow-Headers", "Authorization,Content-Length,X-CSRF-Token,Token,session,X_Requested-With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Pargma")
			context.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Accedss-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//處理請求
		context.Next()
	}
}
