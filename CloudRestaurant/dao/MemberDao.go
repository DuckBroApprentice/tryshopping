// dao操作數據庫
package dao

import (
	"fmt"
	"log"

	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/model"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
)

type MemberDao struct {
	*tool.Orm
}

// 驗證手機號和驗證碼是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode

	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}
	return &sms
}

func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member
}

// 新用戶的數據庫插入操作
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
