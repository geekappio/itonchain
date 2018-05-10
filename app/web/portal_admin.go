package web

import (
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/dal/entity"
)

/**
 * 后台用户登录验证
 */
func HandleAdminLogin(request *model.AdminUser) (*model.ResponseModel){
	util.LogInfo(request)
	return service.UserLogin(request)
}

/**
 * 查询记录行数，总计数据；用于前端计算分页
 */
func HandleGetArticlePendingCount(request model.BaseRequest) (*model.ResponseModel) {
	count, err := service.GetArticlePendingService().GetArticlePendingCount()
	if err != nil {
		util.LogError("查询新获取的文章失败", err)
		return model.NewFailedResponseModel(enum.DB_ERROR,"查询新获取文章文章失败！")
	}
	return model.NewSuccessResponseModelWithData(count)
}

/**
 * 用于新增文章页面查询，并且支持根据文章的标题进行过滤查询
 */
func HandleGetArticlePendingList(request model.ArticlePendingListRequest) (*model.ResponseModel) {
	articlePendings, err := service.GetArticlePendingService().GetArticlePendingList(request.PageNum, request.PageSize, request.ArticleTitle)
	if err != nil {
		util.LogError("查询新获取的文章列表失败", err)
		return model.NewFailedResponseModel(enum.DB_ERROR,"查询新获取的文章列表失败！")
	}
	response := make([]*model.ArticlePendingModel, len(articlePendings))
	for i, articlePending := range articlePendings {
		response[i] = &model.ArticlePendingModel{
			ArticlePendinId: articlePending.Id,
			ArticleTitle:    articlePending.ArticleTitle,
			ArticleFrom:     articlePending.ArticleFrom,
			ArticleUrl:      articlePending.ArticleUrl,
			InternelFid:     articlePending.InternelFid,
			InternelUrl:     articlePending.InternelUrl,
			InternelSize:    articlePending.InternelSize,
			ArticleKeywords: articlePending.ArticleKeywords,
			GmtCreate:       util.TimeFormat(articlePending.GmtCreate),
			GmtUpdate:       util.TimeFormat(articlePending.GmtUpdate),
		}
	}
	return model.NewSuccessResponseModelWithData(articlePendings)
}