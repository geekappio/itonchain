package dao

import (
	"database/sql"

	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
)

var articleSqlMapper *ArticleSqlMapper

func GetArticleSqlMapper(session *xorm.Session) (articleSqlMapper *ArticleSqlMapper) {
	return &ArticleSqlMapper{session:session}
}

type ArticleSqlMapper struct {
	session *xorm.Session
}

func (sqlMapper *ArticleSqlMapper) runtimeSession(sqlTagName string, args ...interface{}) *xorm.Session {
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (articleSqlMapper *ArticleSqlMapper) UpdateArticleFavorite(articleId int64, favoriteTimes int32) (affected sql.Result, err error) {
	paramMap := map[string]interface{}{"id": articleId, "favoriteTimes": favoriteTimes}
	return articleSqlMapper.runtimeSession("update_article_favorite", &paramMap).Execute()
}

func (articleSqlMapper *ArticleSqlMapper) SelectById(articleId int64) (*entity.Article, error) {
	var article entity.Article
	paramMap := map[string]interface{}{"id": articleId}
	_, err := articleSqlMapper.runtimeSession("select_article", &paramMap).Get(&article)
	return &article, err
}
