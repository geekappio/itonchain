package dao

import (
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/xormplus/xorm"
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
		"InternelFid":     articlePending.InternelFid,
		"InternelUrl":     articlePending.InternelUrl,
		"InternelSize":    articlePending.InternelSize,
		"ArticleKeywords": articlePending.ArticleKeywords,
		"GmtCreate":       articlePending.GmtCreate,
		"GmtUpdate":       articlePending.GmtUpdate,
	}
	r, err := self.getSqlTemplateClient("insert_article_pending.stpl", &paramMap).Execute()
	id, _ := r.LastInsertId()
	articlePending.Id = id
	rows, _ := r.RowsAffected()
	return rows, err
}
