package web

import (
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model/field_enum"
)

func HandleUserRegister(request model.WechatUserRequest) (*model.ResponseModel) {
	util.LogInfo(request)
	//TODO 统一校验请求参数
	genderType := field_enum.ValueOf(request.Gender)
	if field_enum.MALE != genderType && field_enum.FEMALE != genderType && field_enum.LADYBOY != genderType {
		return model.NewFailedResponseModel(enum.INVALID_REQUEST_FIELD_VALUE, enum.INVALID_REQUEST_FIELD_VALUE.GetRespMsg())
	}

	userService := service.GetWechatUserService()
	return userService.CreateUser(&request)

}
