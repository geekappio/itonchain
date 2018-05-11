package dao

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

func GetArticleSourceSqlMapper(session *xorm.Session) *ArticleSourceSqlMapper {
	return &ArticleSourceSqlMapper{session: session}
}

type ArticleSourceSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (self *ArticleSourceSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session {
	if self.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return self.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (self *ArticleSourceSqlMapper) CountArticleSources() (int, error) {
	var count []int
	err := self.getSqlTemplateClient("count_article_source.stpl").Find(&count)
	if len(count) == 1 {
		return count[0], err
	} else {
		return 0, err
	}
}

func (self *ArticleSourceSqlMapper) SelectArticleSources(pageNum int, pageSize int) ([]*entity.ArticleSource, error) {
	start := (pageNum - 1) * pageSize
	var articleSources []*entity.ArticleSource
	paramMap := map[string]interface{}{"start":start, "end":pageSize}
	err := self.getSqlTemplateClient("select_article_source_by_page.stpl", &paramMap).Find(&articleSources)
	return articleSources, err
}
