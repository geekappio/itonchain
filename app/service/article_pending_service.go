package service

import (
	"time"

	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
)

type ArticlePendingService struct {
	session *xorm.Session
}

func GetArticlePendingService(session ...*xorm.Session) *ArticlePendingService {
	if len(session) == 0 {
		return &ArticlePendingService{}
	} else {
		return &ArticlePendingService{session: session[0]}
	}
}

func (self *ArticlePendingService) AddArticlePending(title, from, internelFid string, internelUrl string, internelSize int64, keywords string) (bool, error) {
	articlePending := entity.ArticlePending{
		ArticleTitle:    title,
		ArticleFrom:     from,
		InternelFid:     internelFid,
		InternelUrl:     internelUrl,
		InternelSize:    internelSize,
		ArticleKeywords: keywords,
		BaseEntity: entity.BaseEntity{
			GmtCreate: time.Now(),
			GmtUpdate: time.Now(),
		},
	}
	mapper := dao.GetArticlePendingSqlMapper(self.session)

	rows, err := mapper.AddArticlePending(&articlePending)
	if err != nil {
		util.LogError("Error happened when inserting article_pending: ", title, from, internelFid, internelUrl, internelSize, keywords, err)
	}
	return 1 == rows, err
}
