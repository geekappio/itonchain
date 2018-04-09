package web

import (
	. "github.com/geekappio/itonchain/app/enum"
	. "github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/xormplus/xorm"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/dal"
)

func HandleArticleShare(request ArticleShareRequest) (*ArticleShareReturnData, ErrorCode) {
	util.LogInfo(request)
	userService := service.GetWechatUserService()
	user := userService.FindUserByOpenId(request.OpenId)
	if nil == user {
		return nil, SYSTEM_FAILED
	}

	shareService := service.GetArticleShareService()
	ok := shareService.AddArticleShare(user.Id, request.ArticleId)
	if ok {
		times := shareService.CountArticleShare(request.ArticleId)
		return &ArticleShareReturnData{ShareTimes:times}, SYSTEM_SUCCESS
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

	times, respCode := doArticleMark(request, user)
	if respCode.IsSuccess() {
		return &ArticleMarkResponse{MarkTimes:times}, SYSTEM_SUCCESS
	} else {
		return nil, SYSTEM_FAILED
	}
}

func doArticleMark(request ArticleMarkRequest, user *entity.WechatUser) (times int64, code ErrorCode){
	code = dal.Transaction(func(session *xorm.Session) ErrorCode {
		markService := service.GetArticleMarkServiceBySession(session)
		articleService := service.GetArticleServiceBySession(session)
		if MARK.Equals(request.DoMark) {
			err := markService.AddArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if nil != err {
				return DB_INSERT_ERROR
			}
			times, err = articleService.IncMarkTimes(request.ArticleId)
			if nil != err {
				return DB_UPDATE_ERROR
			}
			return SYSTEM_SUCCESS
		} else {
			err := markService.DelArticleMark(user.Id, request.ArticleId, request.CategoryId)
			if nil != err {
				return DB_INSERT_ERROR
			}
			times, err = articleService.DecMarkTimes(request.ArticleId)
			if nil != err {
				return DB_UPDATE_ERROR
			}
			return SYSTEM_SUCCESS
		}
	})
	return
}
