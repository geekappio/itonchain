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
	ok, err := dal.DB.SqlTemplateClient("select_user_by_openId.stpl").Get(wechatUser)
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
	paramMap := map[string]interface{}{"OpenId": wechatUser.OpenId, "NickName": wechatUser.NickName, "AvatarUrl": wechatUser.AvatarUrl, "Gender": wechatUser.Gender, "City": wechatUser.City, "Province": wechatUser.Province, "Country": wechatUser.Country, "Language": wechatUser.Language, "IsDel": wechatUser.IsDel, "GmtCreate": wechatUser.GmtCreate, "GmtUpdate": wechatUser.GmtUpdate};
	result, err := wechatUserSqlMapper.getSqlTemplateClient("insert_wechat_user.stpl", &paramMap).Execute()
	id, _ := result.LastInsertId()
	return id, err
}

// UpdateCategoryOrders call predefined sql template to update category orders
func (wechatUserSqlMapper *WechatUserSqlMapper) UpdateCategoryOrders(openId string, categoryOrders string) (int64, error){
	wechatUser := entity.WechatUser{}
	wechatUser.OpenId = openId
	wechatUser.CategoryOrders = categoryOrders
	return wechatUserSqlMapper.getSqlTemplateClient("update_category_orders_with_openId.stpl").Update(wechatUser)
}
