package model

type ArticleShareRequest struct {
	BaseRequest
	OpenId    string `json:"openId" binding:"required"`
	ArticleId int64    `json:"articleId" binding:"required"`
}

// Return Data field of article share response
type ArticleShareReturnData struct {
	ShareTimes int64 `json:"shareTimes" binding:"required"`
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
	MarkTimes int32 `json:"markTimes"`
}

type ArticleFavoriteRequest struct {
	BaseRequest
	OpenId    string `json:"openId"`
	ArticleId int64    `json:"articleId"`
	DoFavorite string    `json:"doFavorite"`
}

type ArticleListRequest struct {
	BaseRequest
	SearchParams ArticleSearchParams `json:"searchParams"`
	OrderType string `json:"orderType"`
	PageNum int32 `json:"pageNum" binding:"required"`
	PageSize int32 `json:"pageSize"`
}

type ArticleSearchParams struct {
	OpenId string `json:"openId"`
	CategoryId string `json:"categoryId"`
	ArticleTitle string `json:"articleTitle"`
	ArticleLabels string `json:"articleLabels"`
	ArticleKeywords string `json:"articleKeywords"`
	GetTechnology string `json:"getTechnology"`
	GetBlockchain string `json:"getBlockchain"`
	GetRecent string `json:"getRecent"`
	GetMarked string `json:"getMarked"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	State string `json:"state"`
}