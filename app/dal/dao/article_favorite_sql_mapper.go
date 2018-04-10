package dao

import (
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var articleFavoriteSqlMapper *ArticleFavoriteSqlMapper

func GetArticleFavoriteSqlMapper() (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) {
	if articleFavoriteSqlMapper == nil {
		articleFavoriteSqlMapper = &ArticleFavoriteSqlMapper{}
	}

	return articleFavoriteSqlMapper
}

type ArticleFavoriteSqlMapper struct {
}

func (articleFavoriteSqlMapper *ArticleFavoriteSqlMapper) InsertArticleFavorite(articleFavorite *entity.ArticleFavorite) (id int64, err error) {
	return dal.DB.SqlTemplateClient("insert_article_favorite").InsertOne(articleFavorite)
}
