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

func (self *ArticleShareSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session{
	if self.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return self.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (self *ArticleShareSqlMapper) InsertArticleShare(articleShare *entity.ArticleShare) (int64, error) {
	return self.getSqlTemplateClient("insert_article_share").InsertOne(articleShare)
}

// FIXME 测试看是否正确
func (self *ArticleShareSqlMapper) CountArticleShare(articleId int64) (int64, error) {
	var count int
	paramMap := map[string]interface{}{"ArticleId":articleId}
	_, err := self.getSqlTemplateClient("count_article_share", &paramMap).Get(&count)
	return int64(count), err
}