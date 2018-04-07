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
	shareService := service.GetArticleShareService()

	user := entity.WechatUser{
		OpenId:         request.OpenId,
		NickName:       request.NickName,
		AvatarUrl:      request.AvatarlUrl,
		Gender:         request.Gender,
		City:           request.City,
		Province:       request.Province,
		Country:        request.Country,
		Language:       request.Language,
		IsDel:          "NO",
		CategoryOrders: nil,
		BaseEntity : entity.BaseEntity{
			GmtCreate: time.Now(),
			GmtUpdate: time.Now(),
		},
	}

	ok := userService.CreateUser(&user)

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

func UserRegister(c *gin.Context) {
	name := c.Param("name")
	user := entity.WechatUser{
		NickName: name,
	}

	wechatUserService := service.GetWechatUserService()
	ok := wechatUserService.CreateUser(&user)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"returnCode": "1000",
			"returnMsg":  "成功",
			"returnData": user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"returnCode": "1001",
			"returnMsg":  "失败",
			"returnData": user,
		})
	}
}
