package web

import (
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal"

	"strconv"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"github.com/geekappio/itonchain/app/common/common_util"
	"github.com/gin-gonic/gin"
	"github.com/geekappio/itonchain/app/common/logging"
)

func HandleArticleShare(request model.ArticleShareRequest) (*model.ArticleShareReturnData, enum.ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user, err := userService.FindUserByOpenId(request.OpenId)
	if nil != err {
		return nil, enum.SYSTEM_FAILED
	}
	if nil == user {
		return nil, enum.SYSTEM_FAILED
	}

	shareService := service.GetArticleShareService()
	ok, err := shareService.AddArticleShare(user.Id, request.ArticleId)
	if ok && nil == err {
		times, err := shareService.CountArticleShare(request.ArticleId)
		if nil == err {
			return &model.ArticleShareReturnData{ShareTimes: times}, enum.SYSTEM_SUCCESS
		}
	}
	return nil, enum.SYSTEM_FAILED
}

func HandleArticleMark(request model.ArticleMarkRequest) (*model.ArticleMarkResponse, enum.ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user, err := userService.FindUserByOpenId(request.OpenId)
	if nil != err {
		return nil, enum.SYSTEM_FAILED
	}
	if nil == user {
		return nil, enum.SYSTEM_FAILED
	}

	times, respCode := doArticleMark(request, user)
	if respCode.IsSuccess() {
		return &model.ArticleMarkResponse{MarkTimes: times}, enum.SYSTEM_SUCCESS
	} else {
		return nil, enum.SYSTEM_FAILED
	}
}

func doArticleMark(request model.ArticleMarkRequest, user *entity.WechatUser) (times int32, code enum.ErrorCode) {
	code = dal.Transaction(func(session *xorm.Session) enum.ErrorCode {
		markService := service.GetArticleMarkService(session)
		articleService := service.GetArticleService(session)

		markType := field_enum.ValueOf(request.DoMark)
		if field_enum.MARK == markType {
			ok, err := markService.AddArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if !ok && nil != err {
				return enum.DB_INSERT_ERROR
			}
			times, err = articleService.IncMarkTimes(request.ArticleId)
			if nil != err {
				return enum.DB_UPDATE_ERROR
			}
			return enum.SYSTEM_SUCCESS
		} else if field_enum.UNMARK == markType {
			ok, err := markService.DelArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if !ok && nil != err {
				return enum.DB_INSERT_ERROR
			}
			times, err = articleService.DecMarkTimes(request.ArticleId)
			if nil != err {
				return enum.DB_UPDATE_ERROR
			}
			return enum.SYSTEM_SUCCESS
		} else {
			return enum.INVALID_REQUEST_FIELD_VALUE
		}
	})
	return
}

func HandleArticleFavorite(request model.ArticleFavoriteRequest) (*model.ResponseModel) {
	util.LogInfo(request)
	// TODO 校验参数
	favoriteType := field_enum.ValueOf(request.DoFavorite)
	if field_enum.FAVORITE != favoriteType && field_enum.UNFAVORITE != favoriteType {
		return model.NewFailedResponseModel(enum.INVALID_REQUEST_FIELD_VALUE, enum.INVALID_REQUEST_FIELD_VALUE.GetRespMsg())
	}
	userService := service.GetWechatUserService()
	articleService := service.GetArticleService()

	user, err := userService.FindUserByOpenId(request.OpenId)
	if err != nil {
		return model.NewFailedResponseModel(enum.SYSTEM_FAILED, enum.SYSTEM_FAILED.GetRespMsg())
	}
	if nil == user {
		return model.NewFailedResponseModel(enum.USER_NOT_EXISTS, enum.USER_NOT_EXISTS.GetRespMsg())
	}
	// 点赞
	var favoriteTimes int32
	var errUpdate error
	var errorCode enum.ErrorCode
	if field_enum.FAVORITE == favoriteType {
		favoriteTimes, errorCode = doFavorite(request.ArticleId, user.Id, favoriteType)
		if errUpdate != nil {
			return model.NewFailedResponseModel(errorCode, errorCode.GetRespMsg())
		}

	} else {
		favoriteTimes, errUpdate = articleService.UpdateArticleFavorite(request.ArticleId, favoriteType)
		if errUpdate != nil {
			return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR, enum.DB_UPDATE_ERROR.GetRespMsg())
		}
	}
	return &model.ResponseModel{
		ReturnCode: enum.SYSTEM_SUCCESS.GetRespCode(),
		ReturnMsg:  "用户点赞/取消点赞成功",
		ReturnData: favoriteTimes,
	}

}

func doFavorite(articleId int64, userId int64, doFavorite *common_util.EnumType) (favoriteTimes int32, errorCode enum.ErrorCode) {
	errorCode = dal.Transaction(func(session *xorm.Session) enum.ErrorCode {
		articleFavoriteService := service.GetArticleFavoriteService(session)
		_, err := articleFavoriteService.InsertArticleFavorite(articleId, userId)
		articleService := service.GetArticleService(session)
		if err != nil {
			util.LogError(err)
			return enum.DB_INSERT_ERROR
		}
		favoriteTimes, err = articleService.UpdateArticleFavorite(articleId, doFavorite)
		if err != nil {
			util.LogError(err)
			return enum.DB_UPDATE_ERROR
		}
		return enum.SYSTEM_SUCCESS
	})
	return
}

func HandleArticleListQuery(request model.ArticleListRequest) (*model.ResponseModel) {
	openId := request.SearchParams.OpenId
	var user *entity.WechatUser
	if openId != "" {
		wechatUser, err := service.GetWechatUserService().FindUserByOpenId(openId)
		if err != nil {
			return &model.ResponseModel{
				ReturnCode: enum.DB_ERROR.GetRespCode(),
				ReturnMsg:  "查询用户信息失败",
			}
		}
		user = wechatUser
	}

	categoryId := request.SearchParams.CategoryId
	var categoryIdOther int64
	var errOther error
	if categoryId != "" {
		categoryIdOther, errOther = strconv.ParseInt(categoryId, 10, 64)
		if errOther != nil {
			util.LogError("字符串转int64错误", categoryId, errOther)
			return &model.ResponseModel{
				ReturnCode: enum.ILLEGAL_PARAMETERS.GetRespCode(),
				ReturnMsg:  "请求参数不合法",
			}
		}
	}

	var articleIdList *[]int64
	if user != nil || categoryId != "" {
		articleMarkList, err := service.GetArticleMarkService().GetArticleMarkList(user.Id, categoryIdOther)
		if err != nil {
			util.LogError("根据用户id和种类查询文章失败", articleMarkList, err)
			return &model.ResponseModel{
				ReturnCode: enum.DB_ERROR.GetRespCode(),
				ReturnMsg:  "查询文章失败",
			}
		}
		articleIds := make([]int64, len(*articleMarkList))
		for i, value := range *articleMarkList {
			articleIds[i] = value.ArticleId
		}
		articleIdList = &articleIds
	}

	// 分页查询文章
	articleList, err := service.GetArticleService().GetArticleList(request, articleIdList)
	if err != nil {
		util.LogError("根据用户id和种类查询文章失败", articleList, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_ERROR.GetRespCode(),
			ReturnMsg:  "查询文章失败",
		}
	}
	articleListResponse := make([]model.ArticleModel, len(*articleList))
	for i, article := range *articleList {
		articleListResponse[i] = model.ArticleModel{
			Id:              article.Id,
			ArticleTitle:    article.ArticleTitle,
			ArticleFrom:     article.ArticleFrom,
			ArticleUrl:      article.ArticleUrl,
			InternelUrl:     article.InternelUrl,
			ContentType:     article.ContentType,
			Images:          article.Images,
			PreviewLayout:   article.PreviewLayout,
			ArticleLabels:   article.ArticleLabels,
			ArticleKeywords: article.ArticleKeywords,
			FavoriteTimes:   article.FavoriteTimes,
			ViewTimes:       article.ViewTimes,
			MarkTimes:       article.MarkTimes,
			IsTechnology:    article.IsTechnology,
			IsBlockchain:    article.IsBlockchain,
			State:           article.State,
			Comment:         article.Comment,
			GmtCreate:       article.GmtCreate.String(),
		}
	}
	return &model.ResponseModel{
		ReturnCode: enum.SYSTEM_SUCCESS.GetRespCode(),
		ReturnMsg:  "查询文章列表成功",
		ReturnData: articleListResponse,
	}
}

func HandleArticleQuery(request model.ArticleQueryRequest) (*model.ResponseModel) {
	return service.GetArticleService(nil).GetArticle(request);
}

func HandleArticleImageGet((c *gin.Context) {
	logging.Logger.Info("Received request: " + c.Request.RequestURI)

	// values :=c.Request.URL.Query()
	// volume := values.Get("volume")
	// imageKey := values.Get("key")

	// FIXME, 20180505, Henry, Read image from NoSQL storage
	// imageContent :=
	var imageContent []byte
	c.Writer.Write([]byte(imageContent))
	c.Writer.Flush()

	return
}