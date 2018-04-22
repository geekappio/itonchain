package dao

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var articleShareSqlMapper *ArticleShareSqlMapper

func GetArticleShareSqlMapper(session *xorm.Session) *ArticleShareSqlMapper {
	return &ArticleShareSqlMapper{session: session}
}

type ArticleShareSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (self *ArticleShareSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session {
	if self.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return self.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (self *ArticleShareSqlMapper) InsertArticleShare(articleShare *entity.ArticleShare) (int64, error) {
	paramMap := map[string]interface{}{"ArticleId": articleShare.ArticleId, "UserId": articleShare.UserId, "GmtCreate": articleShare.GmtCreate, "GmtUpdate": articleShare.GmtUpdate}
	r, err := self.getSqlTemplateClient("insert_article_share.stpl", &paramMap).Execute()
	id, _ := r.LastInsertId()
	articleShare.Id = id
	rows, _ := r.RowsAffected()
	return rows, err
}

func (self *ArticleShareSqlMapper) CountArticleShare(articleId int64) (int64, error) {
	var count []int64
	paramMap := map[string]interface{}{"ArticleId": articleId}
	err := self.getSqlTemplateClient("count_article_share.stpl", &paramMap).Find(&count)
	if len(count) == 1 {
		return count[0], err
	} else {
		return 0, err
	}
}
