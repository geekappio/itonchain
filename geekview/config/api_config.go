package config

type requestMapping struct {
	UserRegister               string
	ArticleListQuery           string
	ArticleFavorite            string
	ArticleShare               string
	ArticleMark                string
	ArticleCategoryListQuery   string
	ArticleCategoryAdd         string
	ArticleCategoryDelete      string
	ArticleCategoryOrderChange string
	ArticleCategoryInfoChange  string
}

// Export Api request mapping object.
var ApiRequestMapping *requestMapping

func init() {
	ApiRequestMapping = &requestMapping{
		UserRegister:               "/user/register",
		ArticleListQuery:           "/article/list/query",
		ArticleFavorite:            "/article/favorite",
		ArticleShare:               "/article/share",
		ArticleMark:                "/article/mark",
		ArticleCategoryListQuery:   "/article_category/list/query",
		ArticleCategoryAdd:         "/article_category/add",
		ArticleCategoryDelete:      "/article_category/delete",
		ArticleCategoryOrderChange: "/article_category/order/change",
		ArticleCategoryInfoChange:  "/article_category/info/change",
	}
}
