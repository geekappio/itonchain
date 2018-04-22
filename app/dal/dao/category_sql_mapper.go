package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
	"time"
	"github.com/geekappio/itonchain/app/util"
)

var categorySqlMapper *CategorySqlMapper

func GetCategorySqlMapper(session *xorm.Session) (*CategorySqlMapper) {
	return &CategorySqlMapper{session: session}
}

type CategorySqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *CategorySqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session {
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

// AddCategory calls predefined sql template to insert category
func (sqlMapper *CategorySqlMapper) AddCategory(category *entity.Category) (int64, error) {
	paramMap := map[string]interface{}{"UserId": category.UserId, "CategoryName": category.CategoryName, "Description": category.Description, "GmtCreate": time.Now(), "GmtUpdate": time.Now()}
	result, err := sqlMapper.getSqlTemplateClient("insert_category.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	category.Id, _ = result.LastInsertId()
	affectedRows, _ := result.RowsAffected();
	return affectedRows, err
}

// DeleteCategory calls predefined sql template to delete category
func (sqlMapper *CategorySqlMapper) DeleteCategory(categoryId int64, userId int64) (int64, error) {
	paramMap := map[string]interface{}{"CategoryId": categoryId, "UserId": userId}
	result, err := sqlMapper.getSqlTemplateClient("delete_category.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	return result.RowsAffected()
}

// 更新文章种类
func (sqlMapper *CategorySqlMapper) UpdateCategory(category *entity.Category) (int64, error) {
	paramMap := map[string]interface{}{"CategoryName": category.CategoryName, "Description": category.Description, "UserId": category.UserId, "Id": category.Id}
	result, err := sqlMapper.getSqlTemplateClient("update_category.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	return result.RowsAffected()
}

// 查询某个用户的所有category
func (self *CategorySqlMapper) FindByUserId(userId int64) ([]entity.Category, error) {
	var categories []entity.Category
	paramMap := map[string]interface{}{"UserId": userId}
	err := self.getSqlTemplateClient("list_category.stpl", &paramMap).Find(&categories)
	return categories, err
}
