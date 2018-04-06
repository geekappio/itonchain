package web

import (
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
)

func ArticleShareHandler(request model.ArticleShareRequest) (*model.ResponseModel) {
	util.LogInfo(request)
	userService := service.NewWechatUserService()
	shareService := service.NewArticleShareService()

	user := userService.FindUserByOpenId(request.OpenId)
	if nil == user {
		return &model.ResponseModel{ReturnCode: enum.SYSTEM_FAILED}
	}
	articleShare := &entity.ArticleShare{
		ArticleId: request.ArticleId,
		UserId:    user.Id,
	}
	ok := shareService.DoArticleShare(articleShare)
	if ok {
		return &model.ResponseModel{ReturnCode: enum.SYSTEM_SUCCESS, ReturnData: model.ArticleShareReturnData{ShareTimes: 123}}
	} else {
		return &model.ResponseModel{ReturnCode: enum.SYSTEM_FAILED}
	}

}
