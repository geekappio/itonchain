package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
	"time"
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
	paramMap := map[string]interface{}{"ArticleId": articleFavorite.ArticleId, "UserId": articleFavorite.UserId, "GmtCreate": time.Now(), "GmtUPdate": time.Now()}
	result, err := articleFavoriteSqlMapper.getSqlTemplateClient("insert_article_favorite.stpl", &paramMap).Execute()
	articleFavorite.Id, _ = result.LastInsertId();
	affectedRows, _ := result.RowsAffected()
	return affectedRows, err
}
