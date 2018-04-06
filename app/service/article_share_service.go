package service

import (
	"github.com/geekappio/itonchain/app/dal/entity"
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

func (service *ArticleShareService) DoArticleShare(model *entity.ArticleShare) bool {
	println("模拟文件分享完成")
	return true
}
