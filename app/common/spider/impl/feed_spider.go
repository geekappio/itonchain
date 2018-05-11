package impl

import (
	"time"

	"github.com/geekappio/itonchain/app/common/network"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/util"
	goswfsModel "github.com/linxGnu/goseaweedfs/model"
	"github.com/mmcdole/gofeed"
	"github.com/geekappio/itonchain/app/service"
	"sync"
	"github.com/geekappio/itonchain/app/web/api"
)

var (
	signal                = struct{}{}
	feedParser            = gofeed.NewParser()
	articlePendingService = service.GetArticlePendingService()
)

type FeedSpider struct {
	sync.Mutex
	sources  []string
	exitChan chan interface{}
}

func NewFeedSpider() *FeedSpider {
	return &FeedSpider{
		sources:  make([]string, 0),
		exitChan: make(chan interface{}),
	}
}

// FIXME 开启守护进程，目前暂定为每隔 1 小时抓取一次
func (self *FeedSpider) Start() {
	go func() {
		timer := time.NewTimer(0)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				self.capture()
			case <-self.exitChan:
				goto Exit
			}
			timer.Reset(10 * time.Second)
		}
	Exit:
	}()
}

func (self *FeedSpider) Exit() {
	self.exitChan <- signal
}

func (self *FeedSpider) AddSource(source string) bool {
	self.Lock()
	defer self.Unlock()

	// FIXME 时间复杂度为O(n)，简单去重
	for _, s := range self.sources {
		if s == source {
			return false
		}
	}
	self.sources = append(self.sources, source)
	return true
}

func (self *FeedSpider) getSource() []string {
	self.Lock()
	defer self.Unlock()

	snapshot := make([]string, len(self.sources))
	copy(snapshot, self.sources)
	return snapshot
}

func (self *FeedSpider) capture() error {
	for _, source := range self.getSource() {
		// 抓取文章列表并获取最后一篇文章的标记
		lastArticleMark, articles, err := download(source)
		if nil != err {
			return err
		}
		// 遍历并将数据持久化到文件系统后再将记录写入数据库
		for _, article := range articles {
			article.content = localize(article.link, article.content)
			submitResult, err := save(article)
			if nil != err {
				return err
			}
			_, err = articlePendingService.AddArticlePending(
				article.title, article.domain, article.link, submitResult.FileID,
				submitResult.FileURL, submitResult.Size, parseKeywords(article.content))
			if nil != err {
				return err
			}
		}
		// 等执行无误之后再将该文章源最后一篇文章的标记更新
		setLastArticleMark(source, lastArticleMark)
	}
	return nil
}

// 从Feed抓取的文章模型
type FeedArticle struct {
	domain  string
	title   string
	desc    string
	link    string
	pubTime *time.Time
	content string
}

// 将数据持久化到文件系统
func save(feedArticle *FeedArticle) (result *goswfsModel.SubmitResult, err error) {
	name := feedArticle.title
	content := []byte(feedArticle.content)
	return seaweedfs.SubmitRourceContent(name, content, nil)
}

// 下载并解析出文章列表
func download(feedUrl string) (string, []*FeedArticle, error) {
	// 获取该文章源最后一篇文章的标记
	lastArticleMark := getLastArticleMark(feedUrl)

	// 抓取文章原始列表
	feed, err := feedParser.ParseURL(feedUrl)
	if err != nil {
		util.LogWarn("抓取Feed失败：", feedUrl, err)
		return lastArticleMark, nil, err
	}

	articles := make([]*FeedArticle, 0, len(feed.Items))
	// 遍历并转换文章模型
	for _, item := range feed.Items {
		// 如果标记有效且遍历的该次文章与标记匹配则退出遍历
		if "" != lastArticleMark && item.GUID == lastArticleMark {
			break
		}
		// 如果通过Feed抓取不到文章内容，则尝试通过Web再抓取一次
		if "" == item.Content {
			item.Content, _ = network.HttpGet(item.Link)
		}
		_, domain, _ := network.GetUrlInfo(item.Link)
		article := &FeedArticle{
			domain:  domain,
			title:   item.Title,
			desc:    item.Description,
			link:    item.Link,
			pubTime: item.PublishedParsed,
			content: item.Content,
		}
		articles = append(articles, article)
	}
	// 获取最后一篇文章标记
	if len(feed.Items) > 0 {
		lastArticleMark = feed.Items[0].GUID
	}
	return lastArticleMark, articles, nil
}

func getLastArticleMark(feedUrl string) string {
	return redis.Get(api.FEED_LAST_ARTICLE_PREFIX + feedUrl)
}

func setLastArticleMark(feedUrl, GUID string) {
	redis.Set(api.FEED_LAST_ARTICLE_PREFIX+feedUrl, GUID)
}
