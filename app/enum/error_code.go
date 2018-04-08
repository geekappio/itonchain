package enum

const (
	SYSTEM_SUCCESS     = "0000"
	SYSTEM_FAILED      = "0001"
	ILLEGAL_PARAMETERS = "1001"

	DB_ERROR = "2001"
	DB_INSERT_ERROR    = "2002"
	DB_DELETE_ERROR    = "2003"
	DB_UPDATE_ERROR    = "2004"

	USER_NOT_EXISTS      = "3001"
	NULL_CATEGORY_ORDERS = "3002"
)

// FIXME go里边没有类似 enum 的类型，const 只能定义编译期就能确定的类型，暂时没有更好的办法
func GetErrorMsg(errorCode string) string {
	switch errorCode {
	case SYSTEM_SUCCESS:
		return "成功"
	case SYSTEM_FAILED:
		return "系统异常"
	case ILLEGAL_PARAMETERS:
		return "非法参数"
	case DB_ERROR:
		return "数据库操作失败"
	case DB_INSERT_ERROR:
		return "数据库插入失败"
	case DB_DELETE_ERROR:
		return "数据库删除失败"
	case DB_UPDATE_ERROR:
		return "数据库更新失败"

	case USER_NOT_EXISTS:
		return "用户不存在"

	case NULL_CATEGORY_ORDERS:
		return "空的目录顺序项"

	default:
		return "未定义错误码"
	}
}
