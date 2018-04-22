package service

import (
	"time"
	
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
)

type ArticleFavoriteService struct {
	session *xorm.Session
} 

var articleFavoriteService *ArticleFavoriteService

func GetArticleFavoriteService(session ...*xorm.Session) *ArticleFavoriteService {
	if len(session) == 0 {
		return &ArticleFavoriteService{}
	} else {
		return &ArticleFavoriteService{session:session[0]}
	}
}

func (self *ArticleFavoriteService) InsertArticleFavorite(articleId int64, userId int64) (articleFavoriteId int64, err error) {
	articleFavorite := entity.ArticleFavorite{
		ArticleId : articleId,
		UserId : userId,
		BaseEntity : entity.BaseEntity{
			GmtCreate :	time.Now(),
			GmtUpdate :	time.Now(),
	    },
	}
	id, err := dao.GetArticleFavoriteSqlMapper(self.session).InsertArticleFavorite(&articleFavorite)
	if err != nil {
		util.LogError("Error happened when inserting article_favorite: ", articleId, userId, err)
	}
	return id, err
}
