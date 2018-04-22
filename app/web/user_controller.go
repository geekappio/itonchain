package web

import (
	."github.com/geekappio/itonchain/app/model"
	."github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model/field_enum"
)

func HandleUserRegister(request WechatUserRequest) (*ResponseModel) {
	util.LogInfo(request)
	//TODO 统一校验请求参数
	genderType := field_enum.ValueOf(request.Gender)
	if field_enum.MALE != genderType && field_enum.FEMALE != genderType && field_enum.LADYBOY != genderType {
		return &ResponseModel{ReturnCode: INVALID_REQUEST_FIELD_VALUE.GetRespCode(), ReturnMsg:INVALID_REQUEST_FIELD_VALUE.GetRespMsg()}
	}
	userService := service.GetWechatUserService()

	return userService.CreateUser(&request)

}
