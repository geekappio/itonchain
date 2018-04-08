package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var categorySqlMapper *CategorySqlMapper

func GetCategorySqlMapper() (categorySqlMapper *CategorySqlMapper) {
	if categorySqlMapper == nil {
		categorySqlMapper = &CategorySqlMapper{}
	}

	return categorySqlMapper
}

type CategorySqlMapper struct {
}

// Call predefined sql template to insert category
func (sqlMapper *CategorySqlMapper) AddCategory(category *entity.Category) (id int64, err error) {
	return dal.DB.SqlTemplateClient("insert_category").InsertOne(category)
}

// 更新文章种类
func (sqlMapper *CategorySqlMapper) UpdateCategory(category *entity.Category) (id int64 ,err error) {
	return dal.DB.SqlTemplateClient("update_category").Update(category)
}
