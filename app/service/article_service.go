package service

type ArticleService struct {
	
}

var articleService *ArticleService

func GetArticleService() *ArticleService {
	if nil == articleService {
		articleService = &ArticleService{}
	}
	return articleService
}

// TODO 增长并返回mark数
func (self *ArticleService) IncMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}

// TODO 减少并返回mark数
func (self *ArticleService) DecMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}
