package dao

import (
	"database/sql"

	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	."github.com/geekappio/itonchain/app/model"
	"github.com/xormplus/xorm"
)

func GetArticleSqlMapper(session *xorm.Session) (articleSqlMapper *ArticleSqlMapper) {
	return &ArticleSqlMapper{session:session}
}

type ArticleSqlMapper struct {
	common_util.XormSession
	session *xorm.Session
}

func (sqlMapper *ArticleSqlMapper) getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session {
	if sqlMapper.session == nil {
		return dal.DB.SqlTemplateClient(sqlTagName, args ...)
	} else {
		return sqlMapper.session.SqlTemplateClient(sqlTagName, args ...)
	}
}

func (articleSqlMapper *ArticleSqlMapper) UpdateArticleFavorite(articleId int64, favoriteTimes int32) (affected sql.Result, err error) {
	paramMap := map[string]interface{}{"id": articleId, "favoriteTimes": favoriteTimes}
	return articleSqlMapper.getSqlTemplateClient("update_article_favorite.stpl", &paramMap).Execute()
}

func (articleSqlMapper *ArticleSqlMapper) SelectById(articleId int64) (*entity.Article, error) {
	var article entity.Article
	paramMap := map[string]interface{}{"id": articleId}
	_, err := articleSqlMapper.getSqlTemplateClient("select_article", &paramMap).Get(&article)
	return &article, err
}

func (self *ArticleSqlMapper) AddArticleMark(articleId int64, addend int) (int64, error) {
	paramMap := map[string]interface{}{"Id": articleId, "Addend": addend}
	r, err := self.getSqlTemplateClient("add_mark_from_article.stpl", &paramMap).Execute()
	rows, _ := r.RowsAffected()
	return rows, err
}
func (self *ArticleSqlMapper) SelectListByParamsInPage(request ArticleListRequest, articleIdList *[]int64) (*[]entity.Article, error) {
	var articleList []entity.Article
	//TODO 优化分页 类似实现拦截器统一封装
	start := (request.PageNum - 1) * request.PageSize
	end := request.PageNum * request.PageSize
	paramMap := map[string]interface{}{"articleTitle": "%" + request.SearchParams.ArticleTitle + "%", "articleLabels": "%" + request.SearchParams.ArticleLabels + "%",
	"articleKeywords": "%" + request.SearchParams.ArticleKeywords + "%", "getTechnology": request.SearchParams.GetTechnology, "getBlockchain": request.SearchParams.GetBlockchain,
	"getRecent": request.SearchParams.GetRecent, "getMarked": request.SearchParams.GetMarked, "startTime": request.SearchParams.StartTime, "endTime": request.SearchParams.EndTime,
	"state": request.SearchParams.State}
	err := self.getSqlTemplateClient("select_list_by_params_in_page.stpl", &paramMap).In("id",articleIdList).Limit(end, start).Find(&articleList)
	return &articleList, err
}

func (self *ArticleSqlMapper) SelectListByParams(request ArticleListRequest) (*[]entity.Article, error) {
	var articleList []entity.Article
	//TODO 优化分页 类似实现拦截器统一封装
	//start := (request.PageNum - 1) * request.PageSize
	//end := request.PageNum * request.PageSize
	paramMap := map[string]interface{}{"articleTitle": "%" + request.SearchParams.ArticleTitle + "%"}
	err := self.getSqlTemplateClient("select_list_by_params.stpl", &paramMap).Find(&articleList)
	return &articleList, err
}


