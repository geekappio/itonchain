package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
	"time"
	"github.com/geekappio/itonchain/app/util"
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
	paramMap := map[string]interface{}{"OpenId": openId}
	success, err := sqlMapper.getSqlTemplateClient("select_user_by_openId.stpl",&paramMap).Get(wechatUser)
	if success  {
		return wechatUser, err
	} else {
		util.LogError(err)
		return nil, err
	}
}

// InsertUser calls predefined sql template to insert user
func (wechatUserSqlMapper *WechatUserSqlMapper) InsertUser(wechatUser *entity.WechatUser) (int64, error) {
	paramMap := map[string]interface{}{"OpenId": wechatUser.OpenId, "NickName": wechatUser.NickName, "AvatarUrl": wechatUser.AvatarUrl, "Gender": wechatUser.Gender, "City": wechatUser.City, "Province": wechatUser.Province, "Country": wechatUser.Country, "Language": wechatUser.Language, "IsDel": wechatUser.IsDel, "GmtCreate": time.Now(), "GmtUpdate": time.Now()}
	result, err := wechatUserSqlMapper.getSqlTemplateClient("insert_wechat_user.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	wechatUser.Id, _ = result.LastInsertId()
	affectedRows, _ := result.RowsAffected()
	return affectedRows, err
}

// UpdateCategoryOrders call predefined sql template to update category orders
func (wechatUserSqlMapper *WechatUserSqlMapper) UpdateCategoryOrders(openId string, categoryOrders string) (int64, error){
	wechatUser := entity.WechatUser{}
	wechatUser.OpenId = openId
	wechatUser.CategoryOrders = categoryOrders
	paramMap := map[string]interface{}{"OpenId": openId, "CategoryOrders": categoryOrders}
	result, err := wechatUserSqlMapper.getSqlTemplateClient("update_category_orders_with_openId.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	affectedRows, _ := result.RowsAffected()
	return affectedRows, err
}
