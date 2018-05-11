package impl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/geekappio/itonchain/app/common/network"
	"bytes"
	"golang.org/x/net/html"
)

// TODO 从文章内容提取关键字
func parseKeywords(htmlData string) string {
	return ""
}

// 文章在保存之前对其进行本地化操作
func localize(url, htmlData string) string {
	if ("" == url || "" == htmlData) {
		return htmlData
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlData))
	if err != nil {
		return htmlData
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
			s.SetAttr("src", submitResult.FileID)
		}
	})
	buf := bytes.NewBuffer([]byte{})
	rootNode := doc.Nodes[0]	// FIXME
	html.Render(buf, rootNode)
	return buf.String()
}
