package model

type ArticleShareRequest struct {
	BaseRequest
	OpenId    string `json:"openId"`
	ArticleId int64  `json:"articleId"`
}

type ArticleShareReturnData struct {
	BaseResponse
	ShareTimes int32 `json:"shareTimes"`
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
