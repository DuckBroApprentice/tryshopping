// dao操作數據庫
package dao

import (
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/model"
	"github.com/DuckBroApprentice/Shopping/CloudRestaurant/tool"
	"gorm.io/gorm/logger"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
