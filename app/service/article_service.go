package service

import (
	"fmt"
	"strconv"

	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"github.com/geekappio/itonchain/app/util"
	"github.com/jinzhu/copier"
	"github.com/xormplus/xorm"
)

type ArticleService struct {
	session *xorm.Session
}

func GetArticleService(session ...*xorm.Session) *ArticleService {
	if len(session) == 0 {
		return &ArticleService{}
	} else {
		return &ArticleService{session:session[0]}
	}
}

func (self *ArticleService) IncMarkTimes(articleId int64) (int32, error) {
	mapper := dao.GetArticleSqlMapper(self.session)
	_, err := mapper.AddArticleMark(articleId, 1)
	if nil == err {
		var article *entity.Article
		article, err = mapper.SelectById(articleId)
		if nil == err {
			return article.MarkTimes, nil
		}
	}
	return 0, err
}

func (self *ArticleService) DecMarkTimes(articleId int64) (int32, error) {
	mapper := dao.GetArticleSqlMapper(self.session)
	_, err := mapper.AddArticleMark(articleId, -1)
	if nil == err {
		var article *entity.Article
		article, err = mapper.SelectById(articleId)
		if nil == err {
			return article.MarkTimes, nil
		}
	}
	return 0, err
}

func (service *ArticleService) UpdateArticleFavorite(articleId int64, doFavorite *common_util.EnumType) (int32, error) {
	articleSqlMapper := dao.GetArticleSqlMapper(service.session)
	article, err := articleSqlMapper.SelectById(articleId)
	if err != nil {
		util.LogError("查询文章失败", err)
		return 0, err
	}
	var favoriteTimes int32
	if doFavorite == field_enum.FAVORITE {
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

func (service *ArticleService) GetArticleList(request model.ArticleListRequest, articleIdList *[]int64) (*[]entity.Article, error) {
	articleSqlMapper := dao.GetArticleSqlMapper(service.session)
	return articleSqlMapper.SelectListByParamsInPage(request, articleIdList)
}

func (service *ArticleService) GetArticle(request model.ArticleQueryRequest) (*model.ResponseModel) {
	articleModel, err := dao.GetArticleSqlMapper(nil).SelectByArticleIdOrInternelUrl(request.ArticleId, request.InternelFid)

	if err != nil {
		util.LogError("查询文章信息失败", "articleId = "+strconv.FormatInt(request.ArticleId, 10)+";internelUrl = "+request.InternelFid, err)
		return model.NewFailedResponseModel(enum.NOT_FIND_SPECIFIED_ARTICLE, "查询文章信息失败")
	}

	// Get article content
	content, cErr := seaweedfs.DownloadResourceContent(articleModel.InternelFid)
	if cErr != nil {
		util.LogError("查询文章内容失败", "internelFid = ", articleModel.InternelFid,
			", internelUrl = ", articleModel.InternelUrl, cErr)
		return model.NewFailedResponseModel(enum.NOT_FIND_SPECIFIED_ARTICLE, "查询文章内容失败")
	}

	m := model.ArticleModel{}
	copier.Copy(&m, articleModel)
	contentModel := &model.ArticleContentModel{
		ArticleModel: m,
		Content: fmt.Sprintf("%s", content),
	}

	return model.NewSuccessResponseModelWithData(contentModel)
}

