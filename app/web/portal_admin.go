package web

import (
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model/field_enum"
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

/**
 * 将临时文章发布至正式文章表中，并且标记临时文章表状态为已发布，文章表状态为编辑状态
 */
func HandlePublishPengingToArticle(request model.PendingToArticleRequest) (*model.ResponseModel) {
	articlePending, err := service.GetArticlePendingService().GetArticlePending(request.ArticlePendingId)
	if err != nil {
		util.LogError("将Pend文章发布至文章失败", err)
		return model.NewFailedResponseModel(enum.DB_ERROR,"将Pend文章发布至文章失败！")
	}
	if articlePending == nil  {
		util.LogError("Pending文章发布不存在", err)
		return model.NewFailedResponseModel(enum.SYSTEM_FAILED,"Pending文章发布不存在！")
	}
	state := field_enum.ValueOf(articlePending.State)
	if field_enum.ARTICLE_PENDING_PUBLISHED == state {
		util.LogError("Pending文章已经发布，不能再次发布！", err)
		return model.NewFailedResponseModel(enum.SYSTEM_FAILED,"Pending文章已经发布，不能再次发布！")
	}
	// 将临时表中的文章插入至文章表中
	result,err :=service.GetArticleService().AddArticle(articlePending)
	if err != nil || !result {
		util.LogError("将Pend文章发布至文章失败", err)
		return model.NewFailedResponseModel(enum.SYSTEM_FAILED,"将Pend文章发布至文章失败！")
	}
	//并且将临时表中的状态标记为已发布
	resultPend,err :=service.GetArticlePendingService().UpdateArticlePendingStateToPublished(articlePending.Id)
	if err != nil || !resultPend {
		util.LogError("标记临时表文章为发布状态失败", err)
		return model.NewFailedResponseModel(enum.SYSTEM_FAILED,"标记临时表文章为发布状态失败！")
	}
	return model.NewSuccessResponseModel()
}

/**
 * 将文章表中的文章记录标记为上线，即可见
 */
func HandleArticleStateToOnline(request model.ArticleIdsRequest) (*model.ResponseModel) {
	result,err :=service.GetArticleService().UpdateArticleStateToOnline(request.ArticleIds)
	if err != nil || !result {
		util.LogError("文章上线失败", err)
		return model.NewFailedResponseModel(enum.ARTICLE_ONLINE_FAILED,"文章上线失败！")
	}
	return model.NewSuccessResponseModel()
}

/**
 * 将文章表中的文章记录标记为下线，即不可见
 */
func HandleArticleStateToOffline(request model.ArticleIdsRequest) (*model.ResponseModel) {
	result,err :=service.GetArticleService().UpdateArticleStateToOffline(request.ArticleIds)
	if err != nil || !result {
		util.LogError("文章下线失败", err)
		return model.NewFailedResponseModel(enum.ARTICLE_OFFLINE_FAILED,"文章下线失败！")
	}
	return model.NewSuccessResponseModel()
}
