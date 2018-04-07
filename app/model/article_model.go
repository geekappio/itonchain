package model

type ArticleShareRequest struct {
	BaseRequest
	OpenId    string `json:"openId"`
	ArticleId int64    `json:"articleId"`
}

// Return Data field of article share response
type ArticleShareReturnData struct {
	ShareTimes int32 `json:"shareTimes"`
}
