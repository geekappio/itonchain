package model

/**
 * 文章类型修改
 */
type ArticleCategoryChangeRequest struct {
	BaseRequest
	OpenId       string `json:"openId" binding:"required"`
	CategoryId   int    `json:"categoryId" binding:"required"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

// Request of "/article_category/add" api
type ArticleCategoryAddRequest struct {
	BaseRequest
	OpenId       string `json:"openId" binding:"required"`
	CategoryName string `json:"categoryName" binding:"required"`
	Description  string `json:"description"`
	InsertPos    int  `json:"insertPos"`
}

// Return data field of "/article_category/add" api
type ArticleCategoryAddReturnData struct {
	CategoryId int64 `json:"categoryId"`
}

// Request of "/article_category/delete" api
type ArticleCategoryDeleteRequest struct {
	BaseRequest
	OpenId     string `json:"openId" binding:"required"`
	CategoryId int64  `json:"categoryId" binding:"required"`
}

type ArticleCategoryOrderChangeRequest struct {
	BaseRequest
	OpenId     string `json:"openId" binding:"required"`
	CategoryId int64    `json:"categoryId" binding:"required"`
	UpDown     string `json:"upDown" binding:"required"`
}

type ArticleCategoryListRequest struct {
	BaseRequest
	OpenId string `json:"openId"`
}

type ArticleCategoryListResponse struct {
	BaseResponse
	CategoryId   int64  `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	ArticleCount int64  `json:"articleCount"`
	GmtCreate    string `json:"gmtCreate"`
}
