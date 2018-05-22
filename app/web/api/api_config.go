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
	// 后台用户登录，只有geekadmin，密码123456
	AdminLogin				   string
	// 临时文章总数
	ArticlePendingCount		   string
	// 临时文章列表查询
	ArticlePendingListQuery	   string
	// 临时文章保存
	ArticlePendingSave	   string
	// 临时文章publish至生产文章库中
	PublishPengingToArticle	   string
	// 文章上线
	ArticleOnline			   string
	// 文章下线
	ArticleOffline 			   string

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
		AdminLogin:					"/portal/login",
		ArticlePendingCount:		"/portal/article_pending/count",
		ArticlePendingListQuery:	"/portal/article_pending/list",
		ArticlePendingSave:			"/portal/article_pending/save",
		PublishPengingToArticle: 	"/portal/article_pending/publish",
		ArticleOnline:				"/portal/article/online",
		ArticleOffline:				"/portal/article/offline",
	}
}

const (
	RESOURCE_ARTICLE_URI = "/resource/article/"
	RESOURCE_IMAGE_URI   = "/itonchain/resource/image/"
	FEED_LAST_ARTICLE_PREFIX = "FEED_LAST_ARTICLE_PREFIX."
)
