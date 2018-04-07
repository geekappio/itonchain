package enum

const (
	SYSTEM_SUCCESS = "0000"
	SYSTEM_FAILED = "0001"
	ILLEGAL_PARAMETERS = "1001"
	DB_INSERT_ERROR = "2001"
)

// FIXME go里边没有类似 enum 的类型，const 只能定义编译期就能确定的类型，暂时没有更好的办法
func GetErrorMsg(errorCode string) string {
	switch errorCode {
	case SYSTEM_SUCCESS: return "成功"
	case SYSTEM_FAILED: return "系统异常"
	case DB_INSERT_ERROR: return "数据库插入失败"

	default: return "未定义错误码"
	}
}
