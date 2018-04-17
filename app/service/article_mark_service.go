package service

import (
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
	"time"
	"github.com/xormplus/xorm"
)

type ArticleMarkService struct {
	session *xorm.Session
}

func GetArticleMarkService(session ...*xorm.Session) *ArticleMarkService {
	if len(session) == 0 {
		return &ArticleMarkService{}
	} else {
		return &ArticleMarkService{session:session[0]}
	}
}

func (self *ArticleMarkService) AddArticleMark(userId, articleId, categoryId int64) (bool, error) {
	articleMark := entity.ArticleMark{
		ArticleId : articleId,
		UserId : userId,
		CategoryId: categoryId,
		BaseEntity : entity.BaseEntity{
			GmtCreate :	time.Now(),
			GmtUpdate :	time.Now(),
		},
	}
	mapper := dao.GetArticleMarkSqlMapper(self.session)

	mark, err := mapper.SelectArticleMark(userId, articleId, categoryId)
	if nil != mark && nil == err {
		return true, nil
	}

	rows, err := mapper.InsertArticleMark(&articleMark)
	if err != nil {
		util.LogError("Error happened when inserting article_mark: ", articleId, userId, categoryId, err)
	}
	return 1 == rows, err
}

func (self *ArticleMarkService) DelArticleMark(userId, articleId, categoryId int64) (bool, error) {
	return true, dao.GetArticleMarkSqlMapper(self.session).DeleteArticleMark(userId, articleId, categoryId)
}

func (self *ArticleMarkService) GetArticleMarkList(userId int64, categoryId int64) (*[]entity.ArticleMark, error) {
	articleMarkList, err := dao.GetArticleMarkSqlMapper(self.session).SelectArticleMarkListByUserIdAndCategoryId(userId, categoryId)
	if err != nil {
		util.LogError("Error happened when inserting article_mark: ", userId, categoryId, err)
	}
	return articleMarkList, err
}
