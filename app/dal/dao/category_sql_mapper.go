package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"github.com/xormplus/xorm"
)

var categorySqlMapper *CategorySqlMapper

func GetCategorySqlMapper(session *xorm.Session) (*CategorySqlMapper) {
	return &CategorySqlMapper{session:session}
}

type CategorySqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *CategorySqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session{
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

// AddCategory calls predefined sql template to insert category
func (sqlMapper *CategorySqlMapper) AddCategory(category *entity.Category) (int64, error) {
	return sqlMapper.getSqlTemplateClient("insert_category").InsertOne(category)
}

// DeleteCategory calls predefined sql template to delete category
func (sqlMapper *CategorySqlMapper) DeleteCategory(categoryId int64, userId int64) (int64, error) {
	category := entity.Category{}
	category.Id = categoryId
	category.UserId = userId
	category.IsDel = field_enum.NO.Value
	return sqlMapper.getSqlTemplateClient("delete_category").Delete(category)
}

// 更新文章种类
func (sqlMapper *CategorySqlMapper) UpdateCategory(category *entity.Category) (int64, error) {
	return sqlMapper.getSqlTemplateClient("update_category").Update(category)
}

func (self *CategorySqlMapper) FindByUserId(userId int64) ([]entity.Category, error) {
	var categories []entity.Category
	paramMap := map[string]interface{}{"UserId": userId}
	err := self.getSqlTemplateClient("list_category.stpl", &paramMap).Find(&categories)
	return categories, err
}
