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

type ArticleMarkRequest struct {
	BaseRequest
	OpenId     string `json:"openId"`
	ArticleId  int64  `json:"articleId"`
	DoMark     string `json:"doMark"`
	CategoryId int64  `json:"categoryId"`
}

type ArticleMarkResponse struct {
	BaseResponse
	MarkTimes int64 `json:"markTimes"`
}