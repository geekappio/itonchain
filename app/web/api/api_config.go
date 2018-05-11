package api

type requestMapping struct {
	UserRegister               string
	ArticleListQuery           string
	ArticleQuery               string
	ArticleFavorite            string
	ArticleShare               string
	ArticleMark                string
	ArticleCategoryListQuery   string
	ArticleCategoryAdd         string
	ArticleCategoryDelete      string
	ArticleCategoryOrderChange string
	ArticleCategoryInfoChange  string
	ResourceArticleLoad        string
	ResourceImageLoad          string
	WechatPublishAuthen        string
	// 后台服务
	AdminLogin				   string

}

// Export Api request mapping object.
var ApiRequestMapping *requestMapping

func init() {
	ApiRequestMapping = &requestMapping{
		UserRegister:               "/wechat_user/register",
		ArticleQuery:               "/article/query",
		ArticleListQuery:           "/article/list/query",
		ArticleFavorite:            "/article/favorite",
		ArticleShare:               "/article/share",
		ArticleMark:                "/article/mark",
		ArticleCategoryListQuery:   "/article_category/list/query",
		ArticleCategoryAdd:         "/article_category/add",
		ArticleCategoryDelete:      "/article_category/delete",
		ArticleCategoryOrderChange: "/article_category/order/change",
		ArticleCategoryInfoChange:  "/article_category/info/change",
		ResourceArticleLoad:        "/resource/article/:fid",
		ResourceImageLoad:          "/resource/image/:fid",
		WechatPublishAuthen:        "/publish/authentication",
		// 后台服务URL
		AdminLogin:					"admin/login",
	}
}

const (
	RESOURCE_ARTICLE_URI = "/resource/article/"
	RESOURCE_IMAGE_URI   = "/itonchain/resource/image/"
	FEED_LAST_ARTICLE_PREFIX = "FEED_LAST_ARTICLE_PREFIX."
)
