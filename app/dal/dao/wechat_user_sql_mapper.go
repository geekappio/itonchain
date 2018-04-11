package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
)

var wechatUserSqlMapper *WechatUserSqlMapper

func GetWechatUserSqlMapper(session *xorm.Session) (*WechatUserSqlMapper) {
	return &WechatUserSqlMapper{session:session}
}

type WechatUserSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *WechatUserSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session{
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}


// SelectUser calls predefined sql template to insert category
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

// InsertUser calls predefined sql template to insert user
func (wechatUserSqlMapper *WechatUserSqlMapper) InsertUser(wechatUser *entity.WechatUser) (int64, error) {
	return wechatUserSqlMapper.getSqlTemplateClient("insert_wechat_user").InsertOne(wechatUser)
}

// UpdateCategoryOrders call predefined sql template to update category orders
func (wechatUserSqlMapper *WechatUserSqlMapper) UpdateCategoryOrders(openId string, categoryOrders string) (int64, error){
	wechatUser := entity.WechatUser{}
	wechatUser.OpenId = openId
	wechatUser.CategoryOrders = categoryOrders
	return wechatUserSqlMapper.getSqlTemplateClient("update_category_orders_with_openId").Update(wechatUser)
}
