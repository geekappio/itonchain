package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var wechatUserSqlMapper *WechatUserSqlMapper

func GetWechatUserSqlMapper() (wechatUserSqlMapper *WechatUserSqlMapper) {
	if wechatUserSqlMapper == nil {
		wechatUserSqlMapper = &WechatUserSqlMapper{}
	}

	return wechatUserSqlMapper
}

type WechatUserSqlMapper struct {
}

// Call predefined sql template to insert category
func (wechatUserSqlMapper *WechatUserSqlMapper) UserRegister(wechatUser *entity.WechatUser) (bool bool, err error) {
	paramMap := map[string]interface{}{"open_id": wechatUser.OpenId}
	var user entity.WechatUser
	return dal.DB.SqlTemplateClient("select_user_by_openId",paramMap).Get(&user)
}

func (wechatUserSqlMapper *WechatUserSqlMapper) InsertUser(wechatUser *entity.WechatUser) (id int64, err error) {
	return dal.DB.SqlTemplateClient("insert_wechat_user").InsertOne(wechatUser)
}

