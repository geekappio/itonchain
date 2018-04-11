package dao

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
)

var articleMarkSqlMapper *ArticleMarkSqlMapper

func GetArticleMarkSqlMapper(session *xorm.Session) *ArticleMarkSqlMapper {
	return &ArticleMarkSqlMapper{session: session}
}

type ArticleMarkSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (self *ArticleMarkSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session{
	if self.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return self.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (self *ArticleMarkSqlMapper) InsertArticleMark(articleMark *entity.ArticleMark) (int64, error) {
	paramMap := map[string]interface{}{"ArticleId":articleMark.ArticleId, "UserId":articleMark.UserId, "CategoryId":articleMark.CategoryId, "GmtCreate":articleMark.GmtCreate, "GmtUpdate":articleMark.GmtUpdate}
	r, err := self.getSqlTemplateClient("insert_article_mark.stpl", &paramMap).Execute()
	id, _ := r.LastInsertId()
	articleMark.Id = id
	rows, _ := r.RowsAffected()
	return rows, err
}

func (self *ArticleMarkSqlMapper) SelectArticleMarkById(articleMarkId int64) (*entity.ArticleMark, error) {
	var articleMarks []entity.ArticleMark
	paramMap := map[string]interface{}{"Id":articleMarkId}
	err := self.getSqlTemplateClient("select_article_mark_by_id.stpl", &paramMap).Find(&articleMarks)
	if len(articleMarks) == 1 {
		return &articleMarks[0], err
	} else {
		return nil, err
	}
}

func (self *ArticleMarkSqlMapper) SelectArticleMark(userId, articleId, categoryId int64) (*entity.ArticleMark, error) {
	var articleMarks []entity.ArticleMark
	paramMap := map[string]interface{}{"UserId":userId, "ArticleId":articleId, "CategoryId":categoryId}
	err := self.getSqlTemplateClient("select_article_mark.stpl", &paramMap).Find(&articleMarks)
	if len(articleMarks) == 1 {
		return &articleMarks[0], err
	} else {
		return nil, err
	}
}

func (self *ArticleMarkSqlMapper) DeleteArticleMark(userId, articleId, categoryId int64) error {
	paramMap := map[string]interface{}{"UserId":userId, "ArticleId":articleId, "CategoryId":categoryId}
	_, err := self.getSqlTemplateClient("delete_article_mark.stpl", &paramMap).Execute()
	return err
}
