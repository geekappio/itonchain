package web

import (
	"github.com/geekappio/itonchain/app/common/model/api"
	"github.com/geekappio/itonchain/app/common/model/enum"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/util"
	"github.com/gin-gonic/gin"
)

/**
 * 修改文章类别信息
 */
func ArticleCategoryChange(c *gin.Context) {

	openId := c.Param("openId")
	categoryId := c.GetInt("categoryId")
	categoryName := c.Param("categoryName")
	description := c.Param("description")

	// 校验参数
	if openId == "" && categoryId == 0 {
		BuildResopone(c, enum.ILLEGAL_PARAMETERS, "参数非法")
	}

	request := api.ArticleCategoryChange{OpenId: openId, CategoryId: categoryId, CategoryName: categoryName, Description: description}
	// 输出日志
	util.LogInfo(request)

	// 调用服务
	result := service.ArticleCategoryChangeService(request)
	if result {
		BuildResopone(c, enum.SYSTEM_SUCCESS, "处理成功")
	} else {
		BuildResopone(c, enum.SYSTEM_FAILED, "处理失败")
	}
}
