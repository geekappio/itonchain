package model

// Request of "/article_category/order/change" api
type WechatUserRequest struct {
	OpenId string `json:"openId" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	AvatarlUrl string `json:"avatarlUrl"`
	Gender string `json:"gender" binding:"required"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	Language string `json:"language"`
}

type WechatCodeRequest struct {
	Code string `json:"code" binding:"required"`
}
