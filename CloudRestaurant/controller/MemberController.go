package controller

import (
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/param"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/service"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.OPTIONS("/api/login_sms", mc.smsLogin)
}

// 8090/api/sendcode?phone=1234156789
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//發送驗證碼
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "參數解析失敗")
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		tool.Success(context, "發送成功")
		return
	}
	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "發送失敗",
	})

}

// 手機號+短信 登錄的方法
func (mc *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParma
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "參數解析失敗")
		return
	}

	//完成手機+驗證碼登錄
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登錄失敗")
}
