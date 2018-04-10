package service

import (
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model"
)

type ArticleService struct {
	session *xorm.Session
}

func GetArticleService() *ArticleService {
	return &ArticleService{}
}

func GetArticleServiceBySession(session *xorm.Session) *ArticleService {
	return &ArticleService{session:session,}
}

// TODO 增长并返回mark数
func (self *ArticleService) IncMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}

// TODO 减少并返回mark数
func (self *ArticleService) DecMarkTimes(articleId int64) (int64, error) {
	return 0, nil
}

func (service *ArticleService) UpdateArticleFavorite(articleId int64, doFavorite string) (int32, error) {
	articleSqlMapper := dao.GetArticleSqlMapper(nil)
	article, err := articleSqlMapper.SelectById(articleId)
	if err != nil {
		util.LogError("查询文章失败", err)
		return 0, err
	}
	var favoriteTimes int32
	if doFavorite == "FAVORITE" {
		favoriteTimes = article.FavoriteTimes + 1
	} else {
		favoriteTimes = article.FavoriteTimes - 1
	}
	_, errUpdate := articleSqlMapper.UpdateArticleFavorite(articleId, favoriteTimes)
	if errUpdate != nil {
		util.LogError("更新文章点赞数失败", errUpdate)
		return 0, errUpdate
	}
	return favoriteTimes, nil
}

func (service *ArticleService) GetArticleList(request model.ArticleListRequest)  {

}
