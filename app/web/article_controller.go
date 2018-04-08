package web

import (
	. "github.com/geekappio/itonchain/app/enum"
	. "github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
)

func HandleArticleShare(request ArticleShareRequest) (*ArticleShareReturnData, ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user := userService.FindUserByOpenId(request.OpenId)
	if nil == user {
		return nil, SYSTEM_FAILED
	}

	shareService := service.GetArticleShareService()
	ok := shareService.DoArticleShare(user.Id, request.ArticleId)
	if ok {
		return &ArticleShareReturnData{ShareTimes:123}, SYSTEM_SUCCESS
	} else {
		return nil, SYSTEM_FAILED
	}
}

func HandlerArticleMark(request ArticleMarkRequest) (*ArticleMarkResponse, ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user := userService.FindUserByOpenId(request.OpenId)
	if nil == user {
		return nil, SYSTEM_FAILED
	}

	markService := service.GetArticleMarkService()
	articleService := service.GetArticleService()

	var times int64
	// FIXME doMark 参数待定
	if util.EqualsIgnoreCase(request.DoMark, "YES") {
		// TODO 增加事务处理
		_ = markService.AddArticleMark(user.Id, request.ArticleId, request.CategoryId)
		times, _ = articleService.IncMarkTimes(request.ArticleId)
	} else {
		_ = markService.DelArticleMark(user.Id, request.ArticleId, request.CategoryId)
		times, _ = articleService.DecMarkTimes(request.ArticleId)
	}
	return &ArticleMarkResponse{MarkTimes:times}, SYSTEM_SUCCESS
}
