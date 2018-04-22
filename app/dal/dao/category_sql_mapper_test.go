package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/common/logging"
	"testing"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	dal.InitDataSource()
}

func TestAddCategory(t *testing.T) {
	category := &entity.Category{}
	category.UserId = 100
	category.CategoryName = "Test100"
	category.Description = "jdslfjsldfjslkdjflsdjflsjdflksjdfl了圣诞节福利数据库的"

	affectedRows, err := GetCategorySqlMapper(nil).AddCategory(category)
	assert.Equal(t, int64(1), affectedRows)
	assert.Nil(t, err)
}


func TestDeleteCategory(t *testing.T) {
	affectedRows, err := GetCategorySqlMapper(nil).DeleteCategory(9, 100)
	assert.Equal(t, int64(1), affectedRows)
	assert.Nil(t, err)
}

func TestUpdateCategory(t *testing.T) {
	category := &entity.Category{}
	category.Id = 5;
	category.UserId = 100
	category.CategoryName = "Test1001"
	category.Description = "jdslfjsldfjslkdjflsdjflsjdflksjdfl了圣诞节福利数据库的sdfsdfssdfsdf"
	affectedRows, err := GetCategorySqlMapper(nil).UpdateCategory(category)
	assert.Equal(t, int64(1), affectedRows)
	assert.Nil(t, err)
}

func TestFindByUserId(t *testing.T) {
	rows, err := GetCategorySqlMapper(nil).FindByUserId(100)
	assert.NotNil(t, rows)
	assert.Nil(t, err)
}