package service

type ArticleMarkService struct {
}

var articleMarkService *ArticleMarkService

func GetArticleMarkService() *ArticleMarkService {
	if nil == articleMarkService {
		articleMarkService = &ArticleMarkService{}
	}
	return articleMarkService
}

// TODO
func (self *ArticleMarkService) AddArticleMark(userId, articleId, categoryId int64) error {
	return nil
}

// TODO
func (self *ArticleMarkService) DelArticleMark(userId, articleId, categoryId int64) error {
	return nil
}
