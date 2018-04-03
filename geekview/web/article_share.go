package web

import (
	"github.com/geekappio/itonchain/util"
	"github.com/geekappio/itonchain/service"
	"github.com/geekappio/itonchain/common/model/api"
	"github.com/geekappio/itonchain/common/model/dal"
	"github.com/geekappio/itonchain/common/model/enum"
)

func ArticleShareHandler(request api.ArticleShareRequest) (*api.ArticleShareResponse, string) {
	util.LogInfo(request)
	userService := service.NewWechatUserService()
	shareService := service.NewArticleShareService()

	user := userService.FindUserByOpenId(request.OpenId)
	if nil == user {
		return nil, enum.SYSTEM_FAILED
	}
	articleShare := &dal.ArticleShare{
		ArticleId:request.ArticleId,
		UserId:user.Id,
	}
	ok := shareService.DoArticleShare(articleShare)
	if ok {
		return &api.ArticleShareResponse{ShareTimes:123}, enum.SYSTEM_SUCCESS
	} else {
		return nil, enum.SYSTEM_FAILED
	}

}
