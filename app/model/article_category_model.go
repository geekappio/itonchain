package model

/**
 * 文章类型修改
 */
type ArticleCategoryChange struct {
	BaseRequest
	OpenId       string `json:"openId"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

// Request of "/article_category/add" api
type ArticleCategoryAddRequest struct {
	BaseRequest
	OpenId       string `json:"openId"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
	InsertPos    int64  `json:"insertPos"`
}

// Return data field of "/article_category/add" api
type ArticleCategoryAddReturnData struct {
	CategoryId int64 `json:"categoryId"`
}

// Request of "/article_category/delete" api
type ArticleCategoryDeleteRequest struct {
	BaseRequest
	OpenId     string `json:"openId"`
	CategoryId int64  `json:"categoryId"`
}
