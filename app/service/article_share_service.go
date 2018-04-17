package service

import (
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
	"time"
)

var articleShareService *ArticleShareService

// GetArticleCategoryService returns ArticleCategory service instance which provides method calls.
func GetArticleShareService() *ArticleShareService {
	if articleShareService == nil {
		articleShareService = &ArticleShareService{}
	}

	return articleShareService
}

// Implementation struct of ArticleCategory to bind functions wi
type ArticleShareService struct {
}

func (service *ArticleShareService) AddArticleShare(userId, articleId int64) (bool, error) {
	articleShare := entity.ArticleShare{
		ArticleId : articleId,
		UserId : userId,
		BaseEntity : entity.BaseEntity{
			GmtCreate :	time.Now(),
			GmtUpdate :	time.Now(),
		},
	}
	ok, err := dao.GetArticleShareSqlMapper(nil).InsertArticleShare(&articleShare)
	if err != nil {
		util.LogError("Error happened when inserting article_share: ", articleId, userId, err)
	}
	return ok, err
}

func (self *ArticleShareService) CountArticleShare(articleId int64) (int64, error) {
	return dao.GetArticleShareSqlMapper(nil).CountArticleShare(articleId)
}
