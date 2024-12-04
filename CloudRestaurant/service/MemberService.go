package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/dao"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/model"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/param"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {
}

func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParma) *model.Member {
	return nil
}

func (ms *MemberService) Sendcode(phone string) bool {
	//1、產生一個驗證碼
	//code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)) 04v，1後面4個0
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//2、調用skd完成發送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//3、接收返回結果，並判斷發送狀態
	//驗證碼發送成功
	if response.Code == "OK" {
		//將驗證碼保存到數據庫中
		smsCode := model.SmsCode{Phone: phone, Code: code, BizID: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := dao.MemberDao{tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}

	return false
}

//將smsCode插入
