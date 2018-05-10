package network

import (
	"net/url"

	"github.com/geekappio/itonchain/app/common/logging"
	"net/http"
	"io/ioutil"
	"strings"
	"regexp"
)

// GetDomain parse domain from url.
func GetUrlInfo(urlLink string) (schema, domain, port string) {
	u, err := url.Parse(urlLink);
	if err != nil {
		logging.Logger.Error(err)
		return "", "", ""
	}

	return u.Scheme, u.Hostname(), u.Port()
}

// parentUrl必须是完整的，curUrl可以是不完整的，转换出完整的curUrl
func GetCompleteURL(parentUrl, curUrl string) string {
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

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Logger.Error(err)
		return "", nil
	}
	return string(body), nil
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
