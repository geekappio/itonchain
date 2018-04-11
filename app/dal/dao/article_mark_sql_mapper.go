package dao

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
)

type ArticleMarkSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}



func GetArticleMarkSqlMapper(session *xorm.Session) *ArticleMarkSqlMapper  {
	return &ArticleMarkSqlMapper{session: session}
}

func (articleMarkSqlMapper *ArticleMarkSqlMapper) GetArticleMarkList()  {

}
