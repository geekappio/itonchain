package service

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
	"time"
)

type ArticlePendingService struct {
	session *xorm.Session
}

func GetArticlePendingService(session ...*xorm.Session) *ArticlePendingService {
	if len(session) == 0 {
		return &ArticlePendingService{}
	} else {
		return &ArticlePendingService{session:session[0]}
	}
}

func (self *ArticlePendingService) AddArticlePending(title, from, url, keywords string) (bool, error) {
	articlePending := entity.ArticlePending{
		ArticleTitle : title,
		ArticleFrom : from,
		ArticleUrl: url,
		ArticleKeywords: keywords,
		BaseEntity : entity.BaseEntity{
			GmtCreate :	time.Now(),
			GmtUpdate :	time.Now(),
		},
	}
	mapper := dao.GetArticlePendingSqlMapper(self.session)

	rows, err := mapper.AddArticlePending(&articlePending)
	if err != nil {
		util.LogError("Error happened when inserting article_pending: ", title, from, url, keywords, err)
	}
	return 1 == rows, err
}
