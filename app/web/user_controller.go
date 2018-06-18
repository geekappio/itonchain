package web

import (
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/enum"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func HandleUserOpenId(request model.WechatCodeRequest) (*model.ResponseModel){
	client := &http.Client{}
	appId := "wxa8c7a4b982effda9"
	secret := "bab44a049d6af52128958e33332b3506"
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + secret + "&js_code=" + request.Code + "&grant_type=authorization_code"
	resp, err := client.Get(url)
	if err != nil {
		util.LogError(err)
		return model.NewFailedResponseModel(enum.HTTP_ERROR,enum.HTTP_ERROR.GetRespMsg())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.LogError(err)
	}
	jsonStr := string(body)
	util.LogInfo(jsonStr)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err == nil {
		util.LogInfo("openid", data["openid"])
		return model.NewSuccessResponseModelWithData(data["openid"])
	} else {
		util.LogError("json str to struct error")
		return model.NewFailedResponseModel(enum.HTTP_ERROR,enum.HTTP_ERROR.GetRespMsg())
	}
}
