package service

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/web/api"
	goswfsModel "github.com/linxGnu/goseaweedfs/model"
	"github.com/mmcdole/gofeed"
)

const FEED_LAST_ARTICLE_PREFIX = "FEED_LAST_ARTICLE_PREFIX."

var feedParser *gofeed.Parser
var articlePendingService *ArticlePendingService

func init() {
	feedParser = gofeed.NewParser()
	articlePendingService = GetArticlePendingService()
}

type Spider interface {
	Capture(sources []*model.ArticleSource) error
}

type FeedSpider struct {
}

func (self *FeedSpider) Capture(sources []*model.ArticleSource) error {
	for _, source := range sources {
		lastArticleMark, articles, err := download(source.Url)
		if nil != err {
			return err
		}
		for _, article := range articles {
			localize(article)
			submitResult, err := save(article)
			if nil != err {
				return err
			}
			_, err = articlePendingService.AddArticlePending(article.title, article.link, submitResult.FileID, submitResult.FileURL, submitResult.Size, parseKeywords(article.content))
			if nil != err {
				return err
			}
		}
		setLastArticleMark(source.Url, lastArticleMark)
	}
	return nil
}

// TODO 从文章内容提取关键字
func parseKeywords(content string) string {
	return ""
}

// 文章在保存之前对其进行本地化操作
func localize(feedArticle *feedArticle) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(feedArticle.content))
	if err != nil {
		return err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			src = getCompleteURL(feedArticle.link, src)
			resp, err := http.Get(src)
			if err != nil {
				return
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			submitResult, err := seaweedfs.SubmitRourceContent(src, data, nil)
			if err != nil {
				return
			}
			s.SetAttr("src", api.RESOURCE_IMAGE_URI + submitResult.FileID)
		}
	})

	content, err := doc.Html()
	if err != nil {
		return err
	}

	feedArticle.content = content
	return nil
}

func getCompleteURL(parentUrl, curUrl string) string {
	if (strings.HasPrefix(curUrl, "http")) {
		return curUrl
	} else if (strings.HasPrefix(curUrl, "//")) {
		end := strings.Index(parentUrl, "//")
		res := parentUrl[0:end] + curUrl
		return res
	} else if (strings.HasPrefix(curUrl, "/")) {
		prefix, _ := getRoot(parentUrl)
		return prefix + curUrl
	} else {
		prefix, _ := getPath(parentUrl)
		return prefix + curUrl
	}
}

func getRoot(url string) (string, error) {
	return regexUrl("^.*?://[^:/]+:?\\d*", url)
}

func getPath(url string) (string, error) {
	return regexUrl("^.*?://[^:/]+:?\\d*.*/", url)
}

func regexUrl(regex, url string) (string, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	root := reg.FindString(url)
	return root, nil
}

// 从Feed抓取的文章模型
type feedArticle struct {
	title   string
	desc    string
	link    string
	pubTime *time.Time
	content string
}

func save(feedArticle *feedArticle) (result *goswfsModel.SubmitResult, err error) {
	name := feedArticle.title
	content := []byte(feedArticle.content)
	return seaweedfs.SubmitRourceContent(name, content, nil)
}

func download(feedUrl string) (string, []*feedArticle, error) {
	lastArticleMark := getLastArticleMark(feedUrl)

	feed, err := feedParser.ParseURL(feedUrl)
	if err != nil {
		util.LogWarn("抓取Feed失败：", feedUrl, err)
		return lastArticleMark, nil, err
	}

	articles := make([]*feedArticle, len(feed.Items))
	for i, item := range feed.Items {
		if "" == lastArticleMark || item.GUID != lastArticleMark {
			articles[i] = &feedArticle{
				title:   item.Title,
				desc:    item.Description,
				link:    item.Link,
				pubTime: item.PublishedParsed,
				content: item.Content,
			}
		}
	}
	if len(feed.Items) > 0 {
		lastArticleMark = feed.Items[0].GUID
	}
	return lastArticleMark, articles, nil
}

func getLastArticleMark(feedUrl string) string {
	return redis.Get(FEED_LAST_ARTICLE_PREFIX + feedUrl)
}

func setLastArticleMark(feedUrl, GUID string) {
	redis.Set(FEED_LAST_ARTICLE_PREFIX+feedUrl, GUID)
}
