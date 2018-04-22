package service

import (
	"strconv"
	"strings"

	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/dal"
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
	user, userErr := dao.GetWechatUserSqlMapper(nil).SelectUser(request.OpenId)
	if userErr != nil {
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, userErr)
	}
	if user == nil {
		util.LogError("Cannot find user by specified open id: ", request.OpenId, userErr)
		return model.NewFailedResponseModel(enum.USER_NOT_EXISTS, "指定的用户不存在")
	}

	category.Id = user.Id

	// Create transaction session
	session := dal.DB.NewSession();
	defer session.Close()

	_, addErr := dao.GetCategorySqlMapper(session).AddCategory(&category)
	if addErr != nil {
		session.Rollback()

		util.LogError("Error happened when inserting category: ", category, addErr)
		return model.NewFailedResponseModel(enum.DB_INSERT_ERROR, "添加category数据失败")
	}

	if util.StringIsBlack(user.CategoryOrders) {
		// There is no category before, just add current as a new one.
		_, updatedErr := dao.GetWechatUserSqlMapper(session).UpdateCategoryOrders(request.OpenId, strconv.FormatInt(category.Id, 10))
		if updatedErr != nil {
			session.Rollback()

			util.LogError("Error happened when updating user with categorOrders, user_id =  ", user.Id, updatedErr)
			return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR, "添加category数据失败")
		}

	} else {
		// Insert category into proper position
		categories := strings.Split(user.CategoryOrders, config.FIELD_CATEGORY_ORDRES_SEPARATER)
		if request.InsertPos == -1 {
			// Append to last
			categories = append(categories, strconv.FormatInt(category.Id, 10))
		} else {
			categories = util.StringArrayInsert(categories, request.InsertPos, strconv.FormatInt(category.Id, 10))
		}

		_, updatedErr := dao.GetWechatUserSqlMapper(session).UpdateCategoryOrders(request.OpenId, strings.Join(categories, config.FIELD_CATEGORY_ORDRES_SEPARATER))
		if updatedErr != nil {
			session.Rollback()

			util.LogError("Error happened when updating user with categorOrders, user_id =  ", user.Id, updatedErr)
			return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR, "添加category数据失败")
		}
	}

	// Commit transaction.
	session.Commit();
	return model.NewSuccessResponseModelWithData(model.ArticleCategoryAddReturnData{CategoryId: category.Id})
}

func (service *ArticleCategoryService) DeleteArticleCategory(request *model.ArticleCategoryDeleteRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	// Get user model by open id.
	user, err := dao.GetWechatUserSqlMapper(nil).SelectUser(request.OpenId)
	if err != nil {
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, err)
		return model.NewFailedResponseModel(enum.DB_INSERT_ERROR, "更新category数据失败")
	}
	if user == nil {
		util.LogError("Cannot find user by specified open id: ", request.OpenId, err)
		return model.NewFailedResponseModel(enum.USER_NOT_EXISTS, "指定的用户不存在")
	}

	session := dal.DB.NewSession();
	defer session.Close()

	_, deleteErr := dao.GetCategorySqlMapper(session).DeleteCategory(request.CategoryId, user.Id)
	if deleteErr != nil {
		util.LogError("Error happened when deleting category: ", request.CategoryId, err)

		session.Rollback();
		return model.NewFailedResponseModel(enum.DB_DELETE_ERROR,
			"删除category数据失败")
	}

	categories := strings.Split(user.CategoryOrders, config.FIELD_CATEGORY_ORDRES_SEPARATER)
	categories = util.StrigArrayRemove(categories, strconv.FormatInt(request.CategoryId, 10));
	if categories != nil {
		_, updatedErr := dao.GetWechatUserSqlMapper(session).UpdateCategoryOrders(request.OpenId, strings.Join(categories, config.FIELD_CATEGORY_ORDRES_SEPARATER))
		if updatedErr != nil {
			session.Rollback()

			util.LogError("Error happened when updating user with categorOrders, user_id =  ", user.Id, updatedErr)
			return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR,
				"添加category数据失败")
		}
	}

	session.Commit()

	return model.NewSuccessResponseModel();
}

/**
  文章类别统一管理服务实现
 */
func (service *ArticleCategoryService) ArticleCategoryChangeService(request *model.ArticleCategoryChangeRequest) (*model.ResponseModel) {
	// Here calls dao method to access database.
	category := entity.Category{}
	copier.Copy(category, request)

	_, err := dao.GetCategorySqlMapper(nil).UpdateCategory(&category)
	if err != nil {
		util.LogError("Error happened when inserting category: ", category, err)
		return model.NewFailedResponseModel(enum.DB_INSERT_ERROR, "更新category数据失败")
	} else {
		return model.NewFailedResponseModel(enum.SYSTEM_SUCCESS, "更新数据成功")
	}
}

// 查询某个用户的目录分类信息列表
// @param userId 用户Id
func (self *ArticleCategoryService) ListCategoryByUserId(userId int64) ([]entity.Category, error) {
	return dao.GetCategorySqlMapper(nil).FindByUserId(userId)
}
