package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/logging"
)

var wechatUserSQLMapper *WechatUserSQLMapper

func GetWechatUserSQLMapper() (categorySqlMapper *WechatUserSQLMapper) {
	if wechatUserSQLMapper == nil {
		wechatUserSQLMapper = &WechatUserSQLMapper{}
	}

	return wechatUserSQLMapper
}

type WechatUserSQLMapper struct {
}

// Call predefined sql template to insert category
func (sqlMapper *WechatUserSQLMapper) SelectUser(openId string) (user *entity.WechatUser, err error) {
	wechatUser := &entity.WechatUser{}
	ok, err := dal.DB.SqlTemplateClient("select_user_by_openId").Get(wechatUser)
	if err != nil {
		logging.LogError(err)
		return nil, err
	}

	if ok {
		return wechatUser, nil
	} else {
		return nil, err
	}
}
