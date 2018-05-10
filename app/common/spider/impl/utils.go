package impl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/web/api"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/geekappio/itonchain/app/common/network"
)

// TODO 从文章内容提取关键字
func parseKeywords(htmlData string) string {
	return ""
}

// 文章在保存之前对其进行本地化操作
func localize(url, htmlData string) (string, error) {
	if ("" == url || "" == htmlData) {
		return htmlData, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlData))
	if err != nil {
		return htmlData, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			src = network.GetCompleteURL(url, src)
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
			s.SetAttr("src", api.RESOURCE_IMAGE_URI+submitResult.FileID)
		}
	})
	content, err := doc.Html()
	if err != nil {
		return htmlData, err
	}
	return content, nil
}

