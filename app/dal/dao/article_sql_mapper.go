package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"database/sql"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var articleSqlMapper *ArticleSqlMapper

func GetArticleSqlMapper() (articleSqlMapper *ArticleSqlMapper) {
	if articleSqlMapper == nil {
		articleSqlMapper = &ArticleSqlMapper{}
	}
	return articleSqlMapper
}

type ArticleSqlMapper struct {
}

func (articleSqlMapper *ArticleSqlMapper) UpdateArticleFavorite(articleId int64, favoriteTimes int32) (affected sql.Result, err error) {
	paramMap := map[string]interface{}{"id": articleId, "favoriteTimes": favoriteTimes}
	return dal.DB.SqlTemplateClient("update_article_favorite", &paramMap).Execute()
}
func (articleSqlMapper *ArticleSqlMapper) SelectById(articleId int64) (*entity.Article, error) {
	var article entity.Article
	paramMap := map[string]interface{}{"id": articleId}
	_, err := dal.DB.SqlTemplateClient("select_article", &paramMap).Get(&article)
	return &article, err
}
