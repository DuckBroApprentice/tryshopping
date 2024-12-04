package controller

import (
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/service"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

// 8090/api/sendcode?phone=1234156789
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//發送驗證碼
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "參數解析失敗",
		})
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "發送成功",
		})
		return
	}
	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "發送失敗",
	})

}
