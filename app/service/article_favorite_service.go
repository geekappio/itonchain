package service

import (
	"time"
	
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
)

type ArticleFavoriteService struct {
	
} 

var articleFavoriteService *ArticleFavoriteService

func GetArticleFavoriteService() *ArticleFavoriteService {
	if articleFavoriteService == nil {
		articleFavoriteService = &ArticleFavoriteService{}
	}
	return articleFavoriteService
}

func (articleFavoriteService *ArticleFavoriteService) InsertArticleFavorite(articleId int64, userId int64) (articleFavoriteId int64, err error) {
	articleFavorite := entity.ArticleFavorite{
		ArticleId : articleId,
		UserId : userId,
		BaseEntity : entity.BaseEntity{
			GmtCreate :	time.Now(),
			GmtUpdate :	time.Now(),
	    },
	}
	id, err := dao.GetArticleFavoriteSqlMapper(nil).InsertArticleFavorite(&articleFavorite)
	if err != nil {
		util.LogError("Error happened when inserting article_favorite: ", articleId, userId, err)
	}
	return id, err
}
