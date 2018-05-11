package dao

import (
	"testing"
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model"
	"time"
)

func init()  {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	dal.InitDataSource()
}

func TestCategoryFindByUserId(t *testing.T) {
	categories, err := GetCategorySqlMapper(nil).FindByUserId(123321)
	if nil != err {
		t.Fail()
	}
	print(categories)
}

func TestArticleMark(t *testing.T) {
	articleMark := &entity.ArticleMark{
		ArticleId:1,
		UserId:1,
		CategoryId:1,
		BaseEntity:entity.BaseEntity{
			GmtCreate:time.Now(),
			GmtUpdate:time.Now(),
		},
	}

	mapper := GetArticleMarkSqlMapper(nil)

	rows, err := mapper.InsertArticleMark(articleMark)
	if nil != err || 1 != rows {
		t.FailNow()
	}

	entity, err := mapper.SelectArticleMarkById(articleMark.Id)
	if nil != err || nil == entity {
		t.FailNow()
	}
	print(entity)

	err = mapper.DeleteArticleMark(articleMark.UserId, articleMark.ArticleId, articleMark.CategoryId)
	if nil != err {
		t.FailNow()
	}
}

func TestArticleShare(t *testing.T) {
	articleShare := &entity.ArticleShare{
		ArticleId:1,
		UserId:1,
		BaseEntity:entity.BaseEntity{
			GmtCreate:time.Now(),
			GmtUpdate:time.Now(),
		},
	}

	mapper := GetArticleShareSqlMapper(nil)

	rows, err := mapper.InsertArticleShare(articleShare)
	if nil != err || 1 != rows {
		t.FailNow()
	}

	count, err := mapper.CountArticleShare(articleShare.ArticleId)
	if nil != err || 1 >= count {
		t.FailNow()
	}
}

func TestArticleSource(t *testing.T) {
	mapper := GetArticleSourceSqlMapper(nil)

	count, err := mapper.CountArticleSources()
	if nil != err || 1 >= count {
		t.FailNow()
	}
}

func TestArticle(t *testing.T) {
	mapper := GetArticleSqlMapper(nil)

	rows, err := mapper.AddArticleMark(1, -10)
	if nil != err || 1 != rows {
		t.FailNow()
	}
}

func TestWechatUserSqlMapper_InsertUser(t *testing.T) {
	mapper := GetWechatUserSqlMapper(nil)
	wechatUser := entity.WechatUser{
		BaseEntity : entity.BaseEntity{
			GmtCreate:time.Now(),
			GmtUpdate:time.Now(),
		},
		OpenId:        "1",
		NickName:       "哈哈",
		Gender:   "1",
		IsDel:          "NO",
	}
	id, err := mapper.InsertUser(&wechatUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestArticleSqlMapper_SelectListByParamsInPage(t *testing.T) {
	mapper := GetArticleSqlMapper(nil)
	request := model.ArticleListRequest{
		SearchParams : model.ArticleSearchParams{
			ArticleTitle : "数学",
		},
	}
	articleList, err := mapper.SelectListByParams(request)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(articleList)
	}
}
