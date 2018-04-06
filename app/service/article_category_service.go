package service

import (
	"time"

	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
)

var articleCategoryService *ArticleCategoryService

// GetArticleCategoryService returns ArticleCategory service instance which provides method calls.
func GetArticleCategoryService() *ArticleCategoryService {
	if articleCategoryService == nil {
		articleCategoryService = &ArticleCategoryService{}
	}

	return articleCategoryService
}


// Implementation struct of ArticleCategory to bind functions wi
type ArticleCategoryService struct {
}

func (service *ArticleCategoryService) AddArticleCategory(requestModel *model.ArticleCategoryAddRequest) (*model.ResponseModel) {
	// Here calls dao method to access database.
	// TODO ...
	return &model.ResponseModel{}
}

func (service *ArticleCategoryService) DeleteArticleCategory(request *model.ArticleCategoryDeleteRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	// TODO ...
	return &model.ResponseModel{}
}

/**
  文章类别统一管理服务实现
 */
func ArticleCategoryChangeService(request model.ArticleCategoryChange) (bool) {
	affected, err := dal.DB.Where("where id = ? and user_id = ?", request.CategoryId, request.OpenId).
		Update("update category set category_name = ?, description=?, gmt_update=?, update_user=? ",
		request.CategoryName, request.Description, time.Now(), request.OpenId)
	if err != nil && affected != 1 {
		util.LogInfo("更新文章类别失败", err)
		return false
	}
	return true
}
