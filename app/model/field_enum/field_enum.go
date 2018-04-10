package field_enum

import "github.com/geekappio/itonchain/app/common/util"

// 排序类型
var ORDER_TYPE_NO_ORDER = util.DefEnumType("NO_ORDER")
var ORDER_TYPE_TIME_DESC = util.DefEnumType("TIME_DESC")

// YES and No
var YES = util.DefEnumType("YES")
var NO = util.DefEnumType("NO")

// 文章状态
var ARTICLE_STATE_OFFLINE = util.DefEnumType("OFFLINE")
var ARTICLE_STATE_ONLINE = util.DefEnumType("ONLINE")
var ARTICLE_STATE_EDIT = util.DefEnumType("DELETED")

// 收藏和取消收藏
var MARK = util.DefEnumType("MARK")
var UNMARK = util.DefEnumType("UNMARK")

// 关注和取消关注
var FAVORITE = util.DefEnumType("FAVORITE")
var UNFAVORITE = util.DefEnumType("UNFAVORITE")
