package web

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

/**
 * 统一构建返回参数，returnData没有object类型定义，后续扩展
 * 待验证 fixme by xuqihua
 */
func BuildResopone(c *gin.Context, code string, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"returnCode": code,
		"returnMsg":  msg,
		"returnData": "",
	})
}
