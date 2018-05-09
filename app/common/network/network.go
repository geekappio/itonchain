package network

import (
	"net/url"

	"github.com/geekappio/itonchain/app/common/logging"
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
