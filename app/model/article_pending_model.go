package model

type ArticlePendingListRequest struct {
	PageRequest
	ArticleTitle string `json:"searchParams"`
}

type ArticlePendingModel struct {
	BaseResponse
	ArticlePendinId   int64  `json:"articlePendinId"`
	ArticleTitle string `json:"articleTitle"`
	ArticleFrom string  `json:"articleFrom"`
	ArticleUrl string `json:"articleUrl"`
	InternelFid string `json:"internelFid"`
	InternelUrl string `json:"internelUrl"`
	InternelSize int64 `json:"internelSize"`
	ArticleKeywords string `json:"articleKeywords"`
	GmtCreate string `json:"gmtCreate"`
	GmtUpdate string `json:"gmtUpdate"`
}
