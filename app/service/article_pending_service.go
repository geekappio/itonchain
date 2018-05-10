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

func (self *ArticlePendingService) AddArticlePending(title, from, url, internelFid, internelUrl string, internelSize int64, keywords string) (bool, error) {
	articlePending := entity.ArticlePending{
		ArticleTitle:    title,
		ArticleFrom:     from,
		ArticleUrl:      url,
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
		util.LogError("Error happened when inserting article_pending: ", title, from, url, internelFid, internelUrl, internelSize, keywords, err)
	}
	return 1 == rows, err
}

/**
  * 获取记录行数
  */
func (self *ArticlePendingService) GetArticlePendingCount() (int, error) {
	mapper := dao.GetArticlePendingSqlMapper(self.session)
	return mapper.CountArticlePendings()
}

/**
  * 分页获取记录行数
  */
func (self *ArticlePendingService) GetArticlePendingList(pageNum int, pageSize int, articleTitle string) ([]entity.ArticlePending, error) {
	// 如果当前页小于0，默认设置为第一页
	if pageNum < 1 {
		pageNum = 1
	}
	// 如果pageSize小于等于0，默认设置页面大小为20
	if pageSize <= 0 {
		pageSize = 20
	}

	mapper := dao.GetArticlePendingSqlMapper(self.session)
	return mapper.SelectArticlePendings(pageNum, pageSize, articleTitle)
}
/**
 * 根据articlePendingId，查询记录
 */
func (self *ArticlePendingService) GetArticlePending(articlePendingId int64) (*entity.ArticlePending, error) {
	mapper := dao.GetArticlePendingSqlMapper(self.session)
	return mapper.SelectArticlePendingById(articlePendingId)
}
