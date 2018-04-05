package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/geekappio/itonchain/app/common/model/dal"
	"github.com/geekappio/itonchain/app/service"
)

func UserRegister(c *gin.Context) {
	name := c.Param("name")
	user := dal.WechatUser{
		NickName:name,
	}

	wechatUserService := service.NewWechatUserService()
	ok := wechatUserService.CreateUser(&user)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"returnCode": "1000",
			"returnMsg": "成功",
			"returnData": user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"returnCode": "1001",
			"returnMsg": "失败",
			"returnData": user,
		})
	}
}

