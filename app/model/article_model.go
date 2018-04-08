package model

type ArticleShareRequest struct {
	BaseRequest
	OpenId    string `json:"openId" binding:"required"`
	ArticleId int64    `json:"articleId" binding:"required"`
}

// Return Data field of article share response
type ArticleShareReturnData struct {
	ShareTimes int32 `json:"shareTimes" binding:"required"`
}
