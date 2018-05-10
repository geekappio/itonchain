package enum

type ErrorCode string

const (
	// System errors
	SYSTEM_SUCCESS ErrorCode = "0000"
	SYSTEM_FAILED  ErrorCode = "0001"

	// Request & Response errors
	ILLEGAL_PARAMETERS          ErrorCode = "1001"
	INVALID_REQUEST_FIELD_VALUE ErrorCode = "1002"

	// DB error
	DB_ERROR             ErrorCode = "2001"
	DB_INSERT_ERROR      ErrorCode = "2002"
	DB_DELETE_ERROR      ErrorCode = "2003"
	DB_UPDATE_ERROR      ErrorCode = "2004"
	DB_TRANSACTION_ERROR ErrorCode = "2005"

	// Business error
	// 30XX User
	USER_NOT_EXISTS     ErrorCode = "3001"
	USER_ALREADY_EXISTS ErrorCode = "3002"
	PASSWORD_ERROR ErrorCode = "3003"

	// 31XX Article
	NOT_FIND_SPECIFIED_ARTICLE ErrorCode = "3101"

	// 32XX Category
	NULL_CATEGORY_ORDERS        ErrorCode = "3201"
	IS_FIRST_CATEGORY           ErrorCode = "3202"
	IS_LAST_CATEGORY            ErrorCode = "3203"
	NOT_FIND_SPECIFIED_CATEGORY ErrorCode = "3204"
)

func (self ErrorCode) IsSuccess() bool {
	return SYSTEM_SUCCESS == self
}

func (self ErrorCode) IsFailed() bool {
	return !self.IsSuccess()
}
func (self ErrorCode) GetRespCode() string {
	return string(self)
}

// FIXME go里边没有类似 enum 的类型，const 只能定义编译期就能确定的类型，暂时没有更好的办法

func (self ErrorCode) GetRespMsg() string {
	switch self {
	case SYSTEM_SUCCESS:
		return "成功"
	case SYSTEM_FAILED:
		return "系统异常"

	case ILLEGAL_PARAMETERS:
		return "非法参数"
	case INVALID_REQUEST_FIELD_VALUE:
		return "请求参数值非法"

	case DB_ERROR:
		return "数据库操作失败"
	case DB_INSERT_ERROR:
		return "数据库插入失败"
	case DB_DELETE_ERROR:
		return "数据库删除失败"
	case DB_UPDATE_ERROR:
		return "数据库更新失败"
	case DB_TRANSACTION_ERROR:
		return "数据库事务失败"

	case USER_NOT_EXISTS:
		return "用户不存在"
	case NULL_CATEGORY_ORDERS:
		return "空的目录顺序项"
	case IS_FIRST_CATEGORY:
		return "是第一个目录"
	case IS_LAST_CATEGORY:
		return "是最后一个目录"
	case NOT_FIND_SPECIFIED_CATEGORY:
		return "没有发现指定的目录"
	case USER_ALREADY_EXISTS:
		return "用户已存在"

	default:
		return "未定义错误码"
	}
}
