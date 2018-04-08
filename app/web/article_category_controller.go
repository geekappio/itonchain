package web

import (
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
)

/**
 * 修改文章类别信息
 */
func HandleArticleCategoryChange(reqeustModel *model.ArticleCategoryChangeRequest) (response *model.ResponseModel){
	// 输出日志
	util.LogInfo(reqeustModel)

	// 调用服务
	service := service.GetArticleCategoryService()
	return service.ArticleCategoryChangeService(reqeustModel)
}

// HandleArticleCategoryAdd handles request of adding article category.
func HandleArticleCategoryAdd(reqeustModel *model.ArticleCategoryAddRequest) (response *model.ResponseModel) {
	// Call service
	service := service.GetArticleCategoryService()
	return service.AddArticleCategory(reqeustModel)
}

// HandleArticleCategoryDelete handles request of deleting article category.
func HandleArticleCategoryDelete(reqeustModel *model.ArticleCategoryDeleteRequest) (response *model.ResponseModel) {
	// Call service
	service := service.GetArticleCategoryService()
	return service.DeleteArticleCategory(reqeustModel)
}

// HandleArticleCategoryOrderChange handles request of changing order of user's article categories.
func HandleArticleCategoryOrderChange(reqeustModel *model.ArticleCategoryOrderChangeRequest) (response *model.ResponseModel) {
	// Call service
	service := service.GetWechatUserService()
	return service.ChangingArticleCategoryOrder(reqeustModel)
}
