// dao操作數據庫
package dao

import (
	"log"

	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/model"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
