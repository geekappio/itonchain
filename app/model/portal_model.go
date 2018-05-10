package model

/**
 * 后台用户登录模型
 */
type AdminUser struct {
	UserName string      `json:"userName" binding:"required"`
	// password MD5
	Password string      `json:"password" binding:"required"`
}

