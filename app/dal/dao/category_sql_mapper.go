package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model/enum"
)

var categorySqlMapper *CategorySqlMapper

func GetCategorySqlMapper() (*CategorySqlMapper) {
	if categorySqlMapper == nil {
		categorySqlMapper = &CategorySqlMapper{}
	}

	return categorySqlMapper
}

type CategorySqlMapper struct {
}

// AddCategory calls predefined sql template to insert category
func (sqlMapper *CategorySqlMapper) AddCategory(category *entity.Category) (int64, error) {
	return dal.DB.SqlTemplateClient("insert_category").InsertOne(category)
}

// DeleteCategory calls predefined sql template to delete category
func (sqlMapper *CategorySqlMapper) DeleteCategory(categoryId int64, userId int64) (int64, error) {
	category := entity.Category{}
	category.Id = categoryId
	category.UserId = userId
	category.IsDel = enum.NO.Value
	return dal.DB.SqlTemplateClient("delete_category").Delete(category)
}

// 更新文章种类
func (sqlMapper *CategorySqlMapper) UpdateCategory(category *entity.Category) (int64, error) {
	return dal.DB.SqlTemplateClient("update_category").Update(category)
}
