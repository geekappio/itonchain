package service

import (
	"time"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/model/field_enum"
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
		// 文章发布状态默认为未发布
		State:			 field_enum.ARTICLE_PENDING_UNPUBLISH.Value,
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

/**
 * 将临时表中的未发布的文章标记为已发布
 */
func (service *ArticlePendingService) UpdateArticlePendingStateToPublished(articlePendingId int64) (bool, error) {
	rowsAffected, err := dao.GetArticlePendingSqlMapper(nil).UpdateArticlePendingToPublished(articlePendingId)
	if err != nil {
		util.LogError("标记临时表文章为发布状态失败: ", err)
	}
	return 1 == rowsAffected, err
}
