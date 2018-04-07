package web

import (
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
)

func HandleUserRegister(request model.WechatUserRequest) (*model.ResponseModel) {
	util.LogInfo(request)
	//统一校验请求参数
	userService := service.GetWechatUserService()

	return userService.CreateUser(&request)

}
