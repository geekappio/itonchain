package model

type ArticleShareRequest struct {
	BaseRequest
	OpenId    string `json:"openId"`
	ArticleId int    `json:"articleId"`
}

// Return Data field of article share response
type ArticleShareReturnData struct {
	ShareTimes int `json:"shareTimes"`
}
