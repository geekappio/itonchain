package web

import (
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
)

/**
 * 后台用户登录验证
 */
func HandleAdminLogin(request *model.AdminUser) (*model.ResponseModel){
	util.LogInfo(request)
	return service.UserLogin(request)
}