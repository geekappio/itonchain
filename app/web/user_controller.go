package web

import (
	"net/http"
	"time"

	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"github.com/gin-gonic/gin"
	"github.com/geekappio/itonchain/app/enum"
)


func UserRegisterHandler(request model.WechatUserRequest) (*model.ResponseModel) {
	util.LogInfo(request)
	//统一校验请求参数
	userService := service.GetWechatUserService()

	return userService.CreateUser(&request)

}


