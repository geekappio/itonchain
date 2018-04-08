package service

import "github.com/xormplus/xorm"

type ArticleService struct {
	session *xorm.Session
}

func GetArticleService() *ArticleService {
	return &ArticleService{}
}

func GetArticleServiceBySession(session *xorm.Session) *ArticleService {
	return &ArticleService{session:session,}
}

// TODO 增长并返回mark数
func (self *ArticleService) IncMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}

// TODO 减少并返回mark数
func (self *ArticleService) DecMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}
