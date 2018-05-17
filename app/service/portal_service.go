package service

import (
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/util"
)

// 后台服务实现

/**
 * 用户登录验证
 */
func UserLogin(userModel *model.AdminUser) *model.ResponseModel  {
	// 获取配置文件中的用户名和密码
	user := config.App.AdminUser

	// 验证用户名
	if userModel.UserName != user.UserName {
		return model.NewFailedResponseModel(enum.USER_NOT_EXISTS,"用户名不正确！")
	}

	// 验证密码
	if userModel.Password != util.MD5Hash(user.Password) {
		return model.NewFailedResponseModel(enum.PASSWORD_ERROR,"密码错误！")
	}

	return model.NewSuccessResponseModel()
}