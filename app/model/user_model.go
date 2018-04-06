package model

// Request of "/article_category/order/change" api
type ArticleCategoryOrderChangeRequest struct {
	BaseRequest
	OpenId     string `json:"openId"`
	CategoryId int32  `json:"categoryName"`
	UpDown     string `json:"upDown"`
}
