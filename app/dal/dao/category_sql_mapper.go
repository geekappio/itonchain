package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var categorySqlMapper *CategorySQLMapper

func GetCategorySQLMapper() (categorySqlMapper *CategorySQLMapper) {
	if categorySqlMapper == nil {
		categorySqlMapper = &CategorySQLMapper{}
	}

	return categorySqlMapper
}

type CategorySQLMapper struct {
}

// Call predefined sql template to insert category
func (sqlMapper *CategorySQLMapper) AddCategory(category *entity.Category) (id int64, err error) {
	return dal.DB.SqlTemplateClient("insert_category").InsertOne(category)
}
