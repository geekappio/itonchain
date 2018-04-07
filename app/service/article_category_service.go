package service

import (
	"time"

	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"github.com/jinzhu/copier"
	"github.com/geekappio/itonchain/app/dal/dao"
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
	category := entity.Category{}
	copier.Copy(category, requestModel)

	id, err := dao.GetCategorySQLMapper().AddCategory(&category)
	if err != nil {
		util.LogError("Error happened when inserting category: ", category, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg: "添加category数据失败",
		}
	 } else {
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg: "添加category数据失败",
			ReturnData: model.ArticleCategoryAddReturnData{CategoryId: id},
		}
	}
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
