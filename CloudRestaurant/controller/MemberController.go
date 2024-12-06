package controller

import (
	"fmt"

	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/param"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/service"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.POST("/api/login_sms", mc.smsLogin)
	//生產驗證碼調用
	engine.GET("/api/captcha", mc.captcha)
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)

	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)
}

// 用戶名+密碼、驗證碼登錄
func (mc *MemberController) nameLogin(context *gin.Context) {
	//1、解析用戶登錄傳遞參數
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "參數解析失敗")
		return
	}
	//2、驗證驗證碼
	validata := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validata {
		tool.Failed(context, "驗證碼不正確,請重新驗證")
		return
	}
	//3、登錄
	//不直接在這裡實現
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "登錄失敗")
}

// 生成驗證碼
func (mc *MemberController) captcha(context *gin.Context) {
	//todo 生成驗證碼，並返回客戶端
	tool.GenerateCaptcha(context)
}

// 驗證驗證碼是否正確
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "解析失敗")
		return
	}
	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("驗證通過")
	} else {
		fmt.Println("驗證失敗")
	}
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
