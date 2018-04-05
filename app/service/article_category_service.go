package service

import (
	"time"

	"github.com/geekappio/itonchain/app/common/model/api"
	"github.com/geekappio/itonchain/app/dao"
	"github.com/geekappio/itonchain/app/util"
)

/**
  文章类别统一管理服务实现
 */
func ArticleCategoryChangeService(request api.ArticleCategoryChange) (bool) {
	affected, err := dao.DB.Where("where id = ? and user_id = ?", request.CategoryId, request.OpenId).
		Update("update category set category_name = ?, description=?, gmt_update=?, update_user=? ",
		request.CategoryName, request.Description, time.Now(), request.OpenId)
	if err != nil && affected != 1 {
		util.LogInfo("更新文章类别失败", err)
		return false
	}
	return true
}
