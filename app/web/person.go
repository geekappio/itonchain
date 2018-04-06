package web

import (
	"net/http"

	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	name := c.Param("name")
	user := entity.WechatUser{
		NickName: name,
	}

	wechatUserService := service.NewWechatUserService()
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
