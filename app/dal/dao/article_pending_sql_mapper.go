package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/util"
)

func GetArticlePendingSqlMapper(session *xorm.Session) (articleSqlMapper *ArticlePendingSqlMapper) {
	return &ArticlePendingSqlMapper{session: session}
}

type ArticlePendingSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *ArticlePendingSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session {
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (self *ArticlePendingSqlMapper) AddArticlePending(articlePending *entity.ArticlePending) (int64, error) {
	paramMap := map[string]interface{}{
		"ArticleTitle":    articlePending.ArticleTitle,
		"ArticleFrom":     articlePending.ArticleFrom,
		"ArticleUrl":      articlePending.ArticleUrl,
		"InternelFid":     articlePending.InternelFid,
		"InternelUrl":     articlePending.InternelUrl,
		"InternelSize":    articlePending.InternelSize,
		"ArticleKeywords": articlePending.ArticleKeywords,
		"GmtCreate":       articlePending.GmtCreate,
		"GmtUpdate":       articlePending.GmtUpdate,
		"State":       	   articlePending.State,
	}
	r, err := self.getSqlTemplateClient("insert_article_pending.stpl", &paramMap).Execute()
	id, _ := r.LastInsertId()
	articlePending.Id = id
	rows, _ := r.RowsAffected()
	return rows, err
}

// 获取pend文章的记录总数
func (self *ArticlePendingSqlMapper) CountArticlePendings() (int, error) {
	var count []int
	err := self.getSqlTemplateClient("count_article_pending.stpl").Find(&count)
	if len(count) == 1 {
		return count[0], err
	} else {
		return 0, err
	}
}

// 获取pend文章的记录，分页查询
func (self *ArticlePendingSqlMapper) SelectArticlePendings(pageNum int, pageSize int, articleTitle string) ([]entity.ArticlePending, error) {
	// 计算记录起始行和结束行
	start := (pageNum - 1) * pageSize
	var articlePendings []entity.ArticlePending
	paramMap := map[string]interface{}{"start":start, "end":pageSize, "articleTitle" : articleTitle}
	err := self.getSqlTemplateClient("select_article_pending.stpl", &paramMap).Find(&articlePendings)
	return articlePendings, err
}

// 根据ID获取详细文章
func (self *ArticlePendingSqlMapper) SelectArticlePendingById(articlePendingId int64) (*entity.ArticlePending, error) {
	var articlePending entity.ArticlePending
	paramMap := map[string]interface{}{"Id":articlePendingId}
	err := self.getSqlTemplateClient("select_article_pending_by_id.stpl", &paramMap).Find(&articlePending)
	return &articlePending, err
}

// 将临时表中的未发布的文章标记为已发布
func (self *ArticlePendingSqlMapper) UpdateArticlePendingToPublished(articlePendingId int64) (int64, error) {
	paramMap := map[string]interface{}{"Id":articlePendingId}
	result,err := self.getSqlTemplateClient("update_article_pending.stpl", &paramMap).Execute()
	if err != nil {
		util.LogError(err)
		return -1, err
	}

	return result.RowsAffected()
}
