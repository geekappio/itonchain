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

// 临时文章发布至文章状态
var ARTICLE_PENDING_UNPUBLISH = common_util.DefEnumType("UNPUBLISH")
var ARTICLE_PENDING_PUBLISHED = common_util.DefEnumType("PUBLISHED")

// 收藏和取消收藏
var MARK = common_util.DefEnumType("MARK")
var UNMARK = common_util.DefEnumType("UNMARK")

// 关注和取消关注
var FAVORITE = common_util.DefEnumType("FAVORITE")
var UNFAVORITE = common_util.DefEnumType("UNFAVORITE")


var UP = common_util.DefEnumType("UP")
var DOWN = common_util.DefEnumType("DOWN")

//性别
var LADYBOY = common_util.DefEnumType("0")
var MALE = common_util.DefEnumType("1")
var FEMALE = common_util.DefEnumType("2")

// 文章内容格式，枚举：HTML、MD、TXT
var ARTICLE_HTML = common_util.DefEnumType("HTML")
var ARTICLE_MD = common_util.DefEnumType("MD")
var ARTICLE_TXT = common_util.DefEnumType("TXT")

// 预览布局，枚举：TXT(纯文字)、PIC-TXT-TB(图文混排上下)、PIC-TXT-LR(图文混排左右)
var PREVIEW_LAYOUT_TXT = common_util.DefEnumType("TXT")
var PREVIEW_LAYOUT_PIC_TXT_TB = common_util.DefEnumType("PIC-TXT-TB")
var PREVIEW_LAYOUT_PIC_TXT_LR = common_util.DefEnumType("PIC-TXT-LR")

// 文章源类型
var FEED = common_util.DefEnumType("FEED")
var WEB = common_util.DefEnumType("WEB")

// 返回枚举值对应的类型
func ValueOf(value string) *common_util.EnumType {
	return common_util.EnumValueOf(value)
}