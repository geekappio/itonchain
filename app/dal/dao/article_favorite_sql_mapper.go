package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
)

var articleFavoriteSqlMapper *ArticleFavoriteSqlMapper

func GetArticleFavoriteSqlMapper(session *xorm.Session) (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) {
	return &ArticleFavoriteSqlMapper{session: session}
}

type ArticleFavoriteSqlMapper struct {
	session *xorm.Session
}

func (sqlMapper *ArticleFavoriteSqlMapper) runtimeSession(sqlTagName string, args ...interface{}) *xorm.Session{
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) InsertArticleFavorite(articleFavorite *entity.ArticleFavorite) (id int64, err error) {
	return articleFavoriteSqlMapper.runtimeSession("insert_article_favorite").InsertOne(articleFavorite)
}
