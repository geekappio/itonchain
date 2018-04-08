package service

import (
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
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
func (service *ArticleCategoryService) AddArticleCategory(request *model.ArticleCategoryAddRequest) (*model.ResponseModel) {
	// Here calls dao method to access database.
	category := entity.Category{}
	copier.Copy(category, request)

	// Get user model by open id.
	userModel, err := dao.GetWechatUserSqlMapper().SelectUser(request.OpenId)
	if err != nil {
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, err)
	}
	if userModel == nil {
		util.LogError("Cannot find user by specified open id: ", request.OpenId, err)
		return &model.ResponseModel{
			ReturnCode: enum.USER_NOT_EXISTS,
			ReturnMsg:  "指定的用户不存在",
		}
	}

	category.Id = userModel.Id

	// TODO, HENRY, 20180409,
	// 这里要做事务处理，添加category和调整wechat_user.category_orders要一起完成

	id, err := dao.GetCategorySqlMapper().AddCategory(&category)
	if err != nil {
		util.LogError("Error happened when inserting category: ", category, err)
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
	// Get user model by open id.
	userModel, err := dao.GetWechatUserSqlMapper().SelectUser(request.OpenId)
	if err != nil {
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg:  "更新category数据失败",
		}
	}
	if userModel == nil {
		util.LogError("Cannot find user by specified open id: ", request.OpenId, err)
		return &model.ResponseModel{
			ReturnCode: enum.USER_NOT_EXISTS,
			ReturnMsg:  "指定的用户不存在",
		}
	}

	// TODO, HENRY, 20180409,
	// 这里要做事务处理，删除category和调整wechat_user.category_orders要一起完成
	_, er := dao.GetCategorySqlMapper().DeleteCategory(request.CategoryId, userModel.Id)
	if er != nil {
		util.LogError("Error happened when deleting category: ", request.CategoryId, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_DELETE_ERROR,
			ReturnMsg:  "删除category数据失败",
		}
	} else {
		return &model.ResponseModel{
			ReturnCode: enum.SYSTEM_SUCCESS,
		}
	}

	return &model.ResponseModel{}
}

/**
  文章类别统一管理服务实现
 */
func (service *ArticleCategoryService) ArticleCategoryChangeService(request *model.ArticleCategoryChangeRequest) (*model.ResponseModel) {
	// Here calls dao method to access database.
	category := entity.Category{}
	copier.Copy(category, request)

	_,err := dao.GetCategorySqlMapper().UpdateCategory(&category)
	if err != nil {
		util.LogError("Error happened when inserting category: ", category, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg:  "更新category数据失败",
		}
	} else {
		return &model.ResponseModel{
			ReturnCode: enum.SYSTEM_SUCCESS,
			ReturnMsg:  "更新数据成功",
		}
	}
}
