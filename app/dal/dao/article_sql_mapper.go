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
	/*	paramMap := map[string]interface{}{"articleTitle": "%" + request.SearchParams.ArticleTitle + "%", "articleLabels": "%" + request.SearchParams.ArticleLabels + "%",
			"articleKeywords": "%" + request.SearchParams.ArticleKeywords + "%", "getTechnology": request.SearchParams.GetTechnology, "getBlockchain": request.SearchParams.GetBlockchain,
			"getRecent": request.SearchParams.GetRecent, "getMarked": request.SearchParams.GetMarked, "startTime": request.SearchParams.StartTime, "endTime": request.SearchParams.EndTime,
			"state": request.SearchParams.State}*/
	articleTitle := request.SearchParams.ArticleTitle
	session := dal.DB.Where("1=?",1)
	if articleTitle != "" {
		session.And("article_title like ?","%"+articleTitle+"%")
	}
	articleLabels := request.SearchParams.ArticleLabels
	if articleLabels != "" {
		session.And("article_labels like ?","%"+articleLabels+"%")
	}
	articleKeywords := request.SearchParams.ArticleKeywords
	if articleKeywords != "" {
		session.And("article_keywords like ?","%"+articleKeywords+"%")
	}
	getTechnology := request.SearchParams.GetTechnology
	if getTechnology != "" {
		session.And("is_technology = ?",getTechnology)
	}
	getBlockchain := request.SearchParams.GetBlockchain
	if getBlockchain != "" {
		session.And("is_blockchain = ?",getBlockchain)
	}
	state := request.SearchParams.State
	if state != "" {
		session.And("state = ?",state)
	}
	startTime := request.SearchParams.StartTime
	endTime := request.SearchParams.EndTime
	if startTime != "" {
		session.And("gmt_create BETWEEN ? AND ?",startTime, endTime)
	}
	if articleIdList != nil {
		session.In("id", articleIdList)
	}
	err := session.Limit(end, start).Find(&articleList)
	/*sql := "article_title like ? AND article_labels like ? AND article_keywords like ? AND is_technology = ? AND is_blockchain = ? AND state = ? AND gmt_create BETWEEN ? AND ?"
	var err error
	if articleIdList != nil {
		err = dal.DB.Where(sql, "%"+articleTitle+"%", "%"+articleLabels+"%", "%"+articleKeywords+"%", getTechnology, getBlockchain, state, startTime, endTime).In("id", articleIdList).Limit(end, start).Find(&articleList)
	} else {
		err = dal.DB.Where(sql, "%"+articleTitle+"%", "%"+articleLabels+"%", "%"+articleKeywords+"%", getTechnology, getBlockchain, state, startTime, endTime).Limit(end, start).Find(&articleList)
	}*/
	//err := self.getSqlTemplateClient("select_list_by_params_in_page.stpl", &paramMap).In("id", articleIdList).Limit(end, start).Find(&articleList)
	return &articleList, err
}

func (self *ArticleSqlMapper) SelectListByParams(request ArticleListRequest) (*[]entity.Article, error) {
	var articleList []entity.Article
	//TODO 优化分页 类似实现拦截器统一封装
	//start := (request.PageNum - 1) * request.PageSize
	//end := request.PageNum * request.PageSize
	//articleIdListOther:= []int64{1,2}
	//paramMap := map[string]interface{}{"articleTitle":request.SearchParams.ArticleTitle}
	//err := dal.DB.SQL("SELECT * FROM article WHERE article_title like ?articleTitle",&paramMap).In("id",articleIdListOther).Limit(5, 1).Find(&articleList)
	//err := dal.DB.Where("article_title like ? and favorite_times = ?","%" + request.SearchParams.ArticleTitle + "%",2).Find(&articleList)
	//params := []interface{}{"%数学%",2}
	err := dal.DB.Where("1=?",1).And("article_title like ?","s").Find(&articleList)
	//err := dal.DB.Where("article_title like ? and favorite_times = ? AND gmt_create BETWEEN ? AND ?",params,"2018-04-04","2018-04-29").In("id",articleIdListOther).Limit(5, 0).Find(&articleList)
	//paramMap := map[string]interface{}{"articleTitle": "%" + request.SearchParams.ArticleTitle + "%","start": 1,"size":2,"id":2}
	//paramMap := map[string]interface{}{"articleTitle": "%" + request.SearchParams.ArticleTitle + "%","id":"'1','2'"}
	//err := self.getSqlTemplateClient("select_list_by_params.stpl", &paramMap).Find(&articleList)
	return &articleList, err
}


