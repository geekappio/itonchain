package service

import "github.com/xormplus/xorm"

type ArticleMarkService struct {
	session *xorm.Session
}

func GetArticleMarkService() *ArticleMarkService {
	return &ArticleMarkService{}
}

func GetArticleMarkServiceBySession(session *xorm.Session) *ArticleMarkService {
	return &ArticleMarkService{session:session}
}

// TODO
func (self *ArticleMarkService) AddArticleMark(userId, articleId, categoryId int64) error {
	return nil
}

// TODO
func (self *ArticleMarkService) DelArticleMark(userId, articleId, categoryId int64) error {
	return nil
}

func (self *ArticleMarkService) GetArticleMarkList(userId int64, articleId int64){

}
