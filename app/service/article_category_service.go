package service

import (
	"time"

	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/logging"
	"github.com/jinzhu/copier"
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

// AddArticleCategory adds a article cateogry into database.
func (service *ArticleCategoryService) AddArticleCategory(requestModel *model.ArticleCategoryAddRequest) (*model.ResponseModel) {
	// Here calls dao method to access database.
	category := entity.Category{}
	copier.Copy(category, requestModel)

	// Get user model by open id.
	userModel, err := dao.GetWechatUserSQLMapper().SelectUser(requestModel.OpenId)
	if err != nil {
		logging.LogError("Error happened when getting user model from wechat_user table with openId: ", requestModel.OpenId, err)
	}
	category.Id = userModel.Id

	id, err := dao.GetCategorySQLMapper().AddCategory(&category)
	if err != nil {
		logging.LogError("Error happened when inserting category: ", category, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg:  "添加category数据失败",
		}
	} else {
		return &model.ResponseModel{
			ReturnCode: enum.SYSTEM_SUCCESS,
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
		logging.LogInfo("更新文章类别失败", err)
		return false
	}
	return true
}
