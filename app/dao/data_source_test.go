package dao

import (
	"testing"
	. "github.com/geekappio/itonchain/app/config"

	"github.com/ian-kent/go-log/log"
	"github.com/geekappio/itonchain/app/common/model/dal"
)

func TestQuery(t *testing.T) {

	// Setup
	result := t.Run("Init config", func(t *testing.T) {
		var err error
		err = InitAppConfig(DEFAULT_CONFIG_PATH)
		if err != nil {
			log.Fatal(err)
		}

		err = InitDataSource()
		if err != nil {
			log.Fatal(err)
		}
	})
	if !result {
		log.Fatal("初始化配置失败")
		return
	}

	sql_1 := "select * from article_share"
	sql_2 := "select * from article_share where id = ?"
	sql_3 := "select * from article_share where id = ?id"
	param3 := map[string]interface{}{"id": 1}

	// region 1, 返回[]map[string][]byte
	results1, err := DB.QueryBytes(sql_1)
	logErr(t, results1, err)

	results1, err = DB.SQL(sql_2, 1).QueryBytes()
	logErr(t, results1, err)

	results1, err = DB.SQL(sql_3, &param3).QueryBytes()
	logErr(t, results1, err)
	// endregion

	// region 2. 返回[]map[string]string
	results2, err := DB.QueryString(sql_1)
	logErr(t, results2, err)

	results2, err = DB.SQL(sql_2, 1).QueryString()
	logErr(t, results2, err)

	results2, err = DB.SQL(sql_3, &param3).QueryString()
	logErr(t, results2, err)
	// endregion

	// region 3. 返回[]map[string]interface{}
	results3, err := DB.QueryInterface(sql_1)
	logErr(t, results3, err)

	results3, err = DB.SQL(sql_2, 1).QueryInterface()
	logErr(t, results3, err)

	results3, err = DB.SQL(sql_3, &param3).QueryInterface()
	logErr(t, results3, err)

	results3, err = DB.SQL(sql_1).Query().List()
	logErr(t, results3, err)

	results3, err = DB.SQL(sql_1).QueryWithDateFormat("20060102").List()
	logErr(t, results3, err)

	count, err := DB.SQL(sql_1).Query().Count()
	logErr(t, count, err)
	// endregion

	// region 4. 配置XML模板
	results4, err := DB.SqlMapClient("sql_1").Query().List()
	logErr(t, results4, err)

	results4, err = DB.SqlMapClient("sql_2", 1).Query().List()
	logErr(t, results4, err)

	results4, err = DB.SqlMapClient("sql_3", &param3).Query().List()
	logErr(t, results4, err)
	// endregion

	// region 5. 配置动态文件
	results5, err := DB.SqlTemplateClient("select.example.stpl").Query().List()
	logErr(t, results5, err)
	// endregion

	// region 6. 填充模型
	var articleShares []dal.ArticleShare
	err = DB.SQL(sql_1).Find(&articleShares)
	logErr(t, articleShares, err)

	err = DB.SQL(sql_2, 1).Find(&articleShares)
	logErr(t, articleShares, err)

	err = DB.SQL(sql_3, &param3).Find(&articleShares)
	logErr(t, articleShares, err)
	// endregion

	// region 7. XML模板 -> 模型
	err = DB.SqlMapClient("sql_1").Find(&articleShares)
	logErr(t, articleShares, err)
	// endregion

	// region 8. 动态文件 -> 模型
	err = DB.SqlTemplateClient("select.example.stpl").Find(&articleShares)
	logErr(t, articleShares, err)
	// endregion

	// region 9. 查询单条
	var articleShare dal.ArticleShare
	has, err := DB.SQL(sql_1).Get(&articleShare)
	if has {
		logErr(t, articleShare, err)
	}
	// endregion

	// region 事物
	session := DB.NewSession()
	defer session.Close()
	err = session.Begin()
	logErr(t, nil, err)

	_, err = session.Exec(sql_1)
	_, err = session.SQL(sql_1).Get(articleShare)
	err = session.SqlMapClient("sql_1").Find(&articleShare)

	if err != nil {
		err = session.Rollback()
	} else {
		err = session.Commit()
	}



}

func logErr(t *testing.T, result interface{}, err error) {
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
