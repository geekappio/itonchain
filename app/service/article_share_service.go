package service

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

func (service *ArticleShareService) AddArticleShare(userId, articleId int64) bool {
	println("模拟文件分享完成")
	return true
}

func (self *ArticleShareService) CountArticleShare(articleId int64) int64 {
	return 0
}
