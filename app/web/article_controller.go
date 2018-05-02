package web

import (
	. "github.com/geekappio/itonchain/app/enum"
	. "github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal"

	"strconv"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"github.com/geekappio/itonchain/app/common/common_util"
)

func HandleArticleShare(request ArticleShareRequest) (*ArticleShareReturnData, ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user, err := userService.FindUserByOpenId(request.OpenId)
	if nil != err {
		return nil, SYSTEM_FAILED
	}
	if nil == user {
		return nil, SYSTEM_FAILED
	}

	shareService := service.GetArticleShareService()
	ok, err := shareService.AddArticleShare(user.Id, request.ArticleId)
	if ok && nil == err {
		times, err := shareService.CountArticleShare(request.ArticleId)
		if nil == err {
			return &ArticleShareReturnData{ShareTimes: times}, SYSTEM_SUCCESS
		}
	}
	return nil, SYSTEM_FAILED
}

func HandlerArticleMark(request ArticleMarkRequest) (*ArticleMarkResponse, ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user, err := userService.FindUserByOpenId(request.OpenId)
	if nil != err {
		return nil, SYSTEM_FAILED
	}
	if nil == user {
		return nil, SYSTEM_FAILED
	}

	times, respCode := doArticleMark(request, user)
	if respCode.IsSuccess() {
		return &ArticleMarkResponse{MarkTimes:times}, SYSTEM_SUCCESS
	} else {
		return nil, SYSTEM_FAILED
	}
}

func doArticleMark(request ArticleMarkRequest, user *entity.WechatUser) (times int32, code ErrorCode){
	code = dal.Transaction(func(session *xorm.Session) ErrorCode {
		markService := service.GetArticleMarkService(session)
		articleService := service.GetArticleService(session)

		markType := field_enum.ValueOf(request.DoMark)
		if field_enum.MARK == markType {
			ok, err := markService.AddArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if !ok && nil != err {
				return DB_INSERT_ERROR
			}
			times, err = articleService.IncMarkTimes(request.ArticleId)
			if nil != err {
				return DB_UPDATE_ERROR
			}
			return SYSTEM_SUCCESS
		} else if field_enum.UNMARK == markType {
			ok, err := markService.DelArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if !ok && nil != err {
				return DB_INSERT_ERROR
			}
			times, err = articleService.DecMarkTimes(request.ArticleId)
			if nil != err {
				return DB_UPDATE_ERROR
			}
			return SYSTEM_SUCCESS
		} else {
			return INVALID_REQUEST_FIELD_VALUE
		}
	})
	return
}

func HandlerArticleFavorite(request ArticleFavoriteRequest) (*ResponseModel) {
	util.LogInfo(request)
	//TODO 校验参数
	favoriteType := field_enum.ValueOf(request.DoFavorite)
	if field_enum.FAVORITE != favoriteType && field_enum.UNFAVORITE != favoriteType {
		return NewFailedResponseModel(INVALID_REQUEST_FIELD_VALUE,INVALID_REQUEST_FIELD_VALUE.GetRespMsg())
	}
	userService := service.GetWechatUserService()
	articleService := service.GetArticleService()

	user,err := userService.FindUserByOpenId(request.OpenId)
	if err != nil {
		return NewFailedResponseModel(SYSTEM_FAILED,SYSTEM_FAILED.GetRespMsg())
	}
	if nil == user {
		return NewFailedResponseModel(USER_NOT_EXISTS,USER_NOT_EXISTS.GetRespMsg())
	}
	//点赞
	var favoriteTimes int32
	var errUpdate error
	var errorCode ErrorCode
	if field_enum.FAVORITE == favoriteType {
		favoriteTimes, errorCode = doFavorite(request.ArticleId, user.Id,favoriteType)
		if errUpdate != nil {
			return NewFailedResponseModel(errorCode,errorCode.GetRespMsg())
		}

	} else {
		favoriteTimes, errUpdate = articleService.UpdateArticleFavorite(request.ArticleId, favoriteType)
		if errUpdate != nil {
			return NewFailedResponseModel(DB_UPDATE_ERROR,DB_UPDATE_ERROR.GetRespMsg())
		}
	}
	return &ResponseModel{
		ReturnCode: SYSTEM_SUCCESS.GetRespCode(),
		ReturnMsg: "用户点赞/取消点赞成功",
		ReturnData: favoriteTimes,
	}

}

func doFavorite(articleId int64, userId int64, doFavorite *common_util.EnumType) (favoriteTimes int32, errorCode ErrorCode) {
	errorCode = dal.Transaction(func(session *xorm.Session) ErrorCode {
		articleFavoriteService := service.GetArticleFavoriteService(session)
		_, err := articleFavoriteService.InsertArticleFavorite(articleId, userId)
		articleService := service.GetArticleService(session)
		if err != nil {
			util.LogError(err)
			return DB_INSERT_ERROR
		}
		favoriteTimes, err = articleService.UpdateArticleFavorite(articleId, doFavorite)
		if err != nil {
			util.LogError(err)
			return DB_UPDATE_ERROR
		}
		return SYSTEM_SUCCESS
	})
	return
}

func HandlerArticleList(request ArticleListRequest) (*ResponseModel) {
	openId := request.SearchParams.OpenId
	var user *entity.WechatUser
	if openId != "" {
		wechatUser, err := service.GetWechatUserService().FindUserByOpenId(openId)
		if err != nil {
			return &ResponseModel{
				ReturnCode: DB_ERROR.GetRespCode(),
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
			return &ResponseModel{
				ReturnCode: ILLEGAL_PARAMETERS.GetRespCode(),
				ReturnMsg:  "请求参数不合法",
			}
		}
	}

	var articleIdList *[]int64
	if user != nil || categoryId != "" {
		articleMarkList, err := service.GetArticleMarkService().GetArticleMarkList(user.Id, categoryIdOther)
		if err != nil {
			util.LogError("根据用户id和种类查询文章失败", articleMarkList, err)
			return &ResponseModel{
				ReturnCode: DB_ERROR.GetRespCode(),
				ReturnMsg:  "查询文章失败",
			}
		}
		articleIds := make([]int64, len(*articleMarkList))
		for i, value := range *articleMarkList {
			articleIds[i] = value.ArticleId
		}
		articleIdList = &articleIds
	}

	//分页查询文章
	articleList, err := service.GetArticleService().GetArticleList(request, articleIdList)
	if err != nil {
		util.LogError("根据用户id和种类查询文章失败", articleList, err)
		return &ResponseModel{
			ReturnCode: DB_ERROR.GetRespCode(),
			ReturnMsg:  "查询文章失败",
		}
	}
	articleListResponse := make([]ArticleListResponse, len(*articleList))
	for i, article := range *articleList {
		articleListResponse[i] = ArticleListResponse{
			Id:              article.Id,
			ArticleTitle:    article.ArticleTitle,
			ArticleFrom:     article.ArticleFrom,
			ArticleUrl:      article.ArticleUrl,
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
	return &ResponseModel{
		ReturnCode: SYSTEM_SUCCESS.GetRespCode(),
		ReturnMsg:  "查询文章列表成功",
		ReturnData: articleListResponse,
	}
}
