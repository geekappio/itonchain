package field_enum

import "github.com/geekappio/itonchain/app/common/common_util"

// 排序类型
var ORDER_TYPE_NO_ORDER = common_util.DefEnumType("NO_ORDER")
var ORDER_TYPE_TIME_DESC = common_util.DefEnumType("TIME_DESC")

// YES and No
var YES = common_util.DefEnumType("YES")
var NO = common_util.DefEnumType("NO")

// 文章状态
var ARTICLE_STATE_OFFLINE = common_util.DefEnumType("OFFLINE")
var ARTICLE_STATE_ONLINE = common_util.DefEnumType("ONLINE")
var ARTICLE_STATE_EDIT = common_util.DefEnumType("DELETED")

// 收藏和取消收藏
var MARK = common_util.DefEnumType("MARK")
var UNMARK = common_util.DefEnumType("UNMARK")

// 关注和取消关注
var FAVORITE = common_util.DefEnumType("FAVORITE")
var UNFAVORITE = common_util.DefEnumType("UNFAVORITE")


var UP = common_util.DefEnumType("UP")
var DOWN = common_util.DefEnumType("DOWN")

// 返回枚举值对应的类型
func ValueOf(value string) *common_util.EnumType {
	return common_util.EnumValueOf(value)
}