package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/util"
)

var wechatUserSqlMapper *WechatUserSqlMapper

func GetWechatUserSqlMapper() (*WechatUserSqlMapper) {
	if wechatUserSqlMapper == nil {
		wechatUserSqlMapper = &WechatUserSqlMapper{}
	}

	return wechatUserSqlMapper
}

type WechatUserSqlMapper struct {
}

// Call predefined sql template to insert category
func (sqlMapper *WechatUserSqlMapper) SelectUser(openId string) (*entity.WechatUser, error) {
	wechatUser := &entity.WechatUser{}
	ok, err := dal.DB.SqlTemplateClient("select_user_by_openId").Get(wechatUser)
	if err != nil {
		util.LogError(err)
		return nil, err
	}

	if ok {
		return wechatUser, nil
	} else {
		return nil, err
	}
}

// Call predefined sql template to insert category
func (wechatUserSqlMapper *WechatUserSqlMapper) UserRegister(wechatUser *entity.WechatUser) (bool, error) {
	paramMap := map[string]interface{}{"open_id": wechatUser.OpenId}
	var user entity.WechatUser
	return dal.DB.SqlTemplateClient("select_user_by_openId",paramMap).Get(&user)
}

func (wechatUserSqlMapper *WechatUserSqlMapper) InsertUser(wechatUser *entity.WechatUser) (int64, error) {
	return dal.DB.SqlTemplateClient("insert_wechat_user").InsertOne(wechatUser)
}
