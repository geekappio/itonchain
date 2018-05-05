package web

import (
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/dal/entity"
)

/**
 * 修改文章类别信息
 */
func HandleArticleCategoryChange(reqeustModel *model.ArticleCategoryChangeRequest) (*model.ResponseModel){
	// 输出日志
	util.LogInfo(reqeustModel)

	// 调用服务
	service := service.GetArticleCategoryService()
	return service.ArticleCategoryChangeService(reqeustModel)
}

// HandleArticleCategoryAdd handles request of adding article category.
func HandleArticleCategoryAdd(reqeustModel *model.ArticleCategoryAddRequest) (*model.ResponseModel) {
	// Call service
	service := service.GetArticleCategoryService()
	return service.AddArticleCategory(reqeustModel)
}

// HandleArticleCategoryDelete handles request of deleting article category.
func HandleArticleCategoryDelete(reqeustModel *model.ArticleCategoryDeleteRequest) (*model.ResponseModel) {
	// Call service
	service := service.GetArticleCategoryService()
	return service.DeleteArticleCategory(reqeustModel)
}

// HandleArticleCategoryOrderChange handles request of changing order of user's article categories.
func HandleArticleCategoryOrderChange(reqeustModel *model.ArticleCategoryOrderChangeRequest) (*model.ResponseModel) {
	// Call service
	service := service.GetWechatUserService()
	return service.ChangingArticleCategoryOrder(reqeustModel)
}

// HandleArticleCategoryListQuery
func HandleArticleCategoryListQuery(request model.ArticleCategoryListRequest) ([]*model.ArticleCategoryModel, enum.ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user, err := userService.FindUserByOpenId(request.OpenId)
	if nil != err {
		return nil, enum.SYSTEM_FAILED
	}
	if nil == user {
		return nil, enum.SYSTEM_FAILED
	}

	categoryService := service.GetArticleCategoryService()
	categories, _ := categoryService.ListCategoryByUserId(user.Id)
	response := make([]*model.ArticleCategoryModel, len(categories))
	for i, c := range categories {
		response[i] = buildCategoryListResponse(&c)
	}
	return response, enum.SYSTEM_SUCCESS
}

func buildCategoryListResponse(category *entity.Category) *model.ArticleCategoryModel {
	return &model.ArticleCategoryModel{
		CategoryId:category.Id,
		CategoryName:category.CategoryName,
		ArticleCount:category.ArticleCount,
		GmtCreate:util.TimeFormat(category.GmtCreate),
	}
}
