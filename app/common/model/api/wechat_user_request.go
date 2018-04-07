package api

type WechatUserRequest struct {
	OpenId string `json:"user" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	AvatarlUrl string `json:"avatarlUrl"`
	Gender string `json:"gender" binding:"required"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	Language string `json:"language"`
}
