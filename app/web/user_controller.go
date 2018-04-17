package web

import (
	."github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
)

func HandleUserRegister(request WechatUserRequest) (*ResponseModel) {
	util.LogInfo(request)
	//TODO 统一校验请求参数
	userService := service.GetWechatUserService()

	return userService.CreateUser(&request)

}
