package api

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
