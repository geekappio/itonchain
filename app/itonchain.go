package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/util"

	"crypto/sha1"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/web"
	"github.com/geekappio/itonchain/app/web/api"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func initConfig() error {

	// Parse command line arguments
	// App file path
	configPath := flag.String("config", config.DEFAULT_CONFIG_PATH, "Needs config file path.")
	flag.Parse()

	var err error

	// Init application configurations.
	err = config.InitAppConfig(*configPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Init Loggers
	err = logging.InitLoggers()
	if err != nil {
		log.Fatal(err)
	}
	// Init redis.
	err = redis.InitRedis()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Init SeaWeedFS
	seaweedfs.InitSeaWeedFS();

	// Init database
	err = dal.InitDataSource()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {

	// Init application configurations.
	err := initConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	gin.SetMode(config.App.RunMode) // 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()         // 获得路由实例
	router.Use(Middleware)
	// 浏览器里显示静态页面
	router.GET("/", rootHandler)

	// 加载模板资源
	router.LoadHTMLGlob("resource/web/template/*")
	// 加载静态资源
	router.Static("/static", "resource/web/static/")

	//router.GET("/login", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "" ,gin.H{
	//		"key":"value",
	//	})
	//})

	// 微信小程序认证
	router.GET(api.ApiRequestMapping.WechatPublishAuthen, authenticateGeekappPublishHandler)

	// Loading article image from NoSQL storage.
	router.GET(api.ApiRequestMapping.ResourceImageLoad, web.HandleResourceGet)

	// Loading article from NoSQL storage.
	router.GET(api.ApiRequestMapping.ResourceArticleLoad, web.HandleResourceGet)

	// 注册用户
	util.AddPostRouter(router, api.ApiRequestMapping.UserRegister, web.HandleUserRegister)

	// 查询/搜索文章列表
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleListQuery, web.HandleArticleListQuery)

	// 查询文章详情
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleQuery, web.HandleArticleQuery)

	// 点赞文章
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleFavorite, web.HandleArticleFavorite)

	// 分享文章
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleShare, web.HandleArticleShare)

	// 收藏文章
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleMark, web.HandleArticleMark)

	// 查询分类列表
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleCategoryListQuery, web.HandleArticleCategoryListQuery)

	// 修改修改文章类别信息
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleCategoryInfoChange, web.HandleArticleCategoryChange)

	// 添加文章类目
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleCategoryAdd, web.HandleArticleCategoryAdd)

	// 删除文章类目
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleCategoryDelete, web.HandleArticleCategoryDelete)

	// 修改文章类目次序
	util.AddPostRouter(router, api.ApiRequestMapping.ArticleCategoryOrderChange, web.HandleArticleCategoryOrderChange)

	/**
	 * 后台服务API
	 */
	/* region 返回Json，需要前端JS获取数据并展示 */
	//// 后台用户登录
	//util.AddPostRouter(router, api.ApiRequestMapping.AdminLogin, web.HandleAdminLogin)
	//// 获取pending的文章总数
	//util.AddPostRouter(router, api.ApiRequestMapping.ArticlePendingCount, web.HandleGetArticlePendingCount)
	//// 获取pending文章列表，支持文件标题进行查询
	//util.AddPostRouter(router, api.ApiRequestMapping.ArticlePendingListQuery, web.HandleGetArticlePendingList)
	//// 将Penging文章推送至文章表（状态为编辑）
	//util.AddPostRouter(router, api.ApiRequestMapping.PublishPengingToArticle, web.HandlePublishPengingToArticle)
	//// 文章上线
	//util.AddPostRouter(router, api.ApiRequestMapping.ArticleOnline, web.HandleArticleStateToOnline)
	//// 文章下线
	//util.AddPostRouter(router, api.ApiRequestMapping.ArticleOffline, web.HandleArticleStateToOffline)
	/* endregion */

	/* region 内部使用模板，返回动态页面 */
	// 获取pending文章列表，支持文件标题进行查询
	router.GET(api.ApiRequestMapping.ArticlePendingListQuery, web.ArticlePendingList)
	/* endregion */

	// Handle websocket
	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	router.Run(config.App.Server.Address)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func rootHandler(c *gin.Context) {
	w := c.Writer
	content, err := ioutil.ReadFile("resource/web/pages/index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func CheckWechatPublishSign(signature string, timestamp string, nonce string) bool {

	arrs := []string{config.App.GeekappPublish.Token, timestamp, nonce}
	sort.Strings(arrs)

	raw := strings.Join(arrs, "")
	h := sha1.New();
	h.Write([]byte(raw))
	sha := fmt.Sprintf("%x", h.Sum(nil))
	return signature == sha
}

func authenticateGeekappPublishHandler(c *gin.Context) {
	logging.Logger.Info("Received request: " + c.Request.RequestURI)

	values := c.Request.URL.Query()
	signature := values.Get("signature")
	timestamp := values.Get("timestamp")
	nonce := values.Get("nonce")
	echostr := values.Get("echostr")

	isValid := CheckWechatPublishSign(signature, timestamp, nonce)

	if isValid {
		logging.Logger.Info("publish.geekapp authentication success.")
		c.Writer.Write([]byte(echostr))
		c.Writer.Flush()
	} else {
		logging.Logger.Info("publish.geekapp authentication failed.")
		c.Writer.Write([]byte("Error"))
		c.Writer.Flush()
	}
}
