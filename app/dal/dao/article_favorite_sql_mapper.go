package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
)

var articleFavoriteSqlMapper *ArticleFavoriteSqlMapper

func GetArticleFavoriteSqlMapper(session *xorm.Session) (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) {
	return &ArticleFavoriteSqlMapper{session: session}
}

type ArticleFavoriteSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *ArticleFavoriteSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session{
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) InsertArticleFavorite(articleFavorite *entity.ArticleFavorite) (id int64, err error) {
	return articleFavoriteSqlMapper.getSqlTemplateClient("insert_article_favorite.stpl").InsertOne(articleFavorite)
}
