package service

import (
	"testing"

	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/model"
)

func init()  {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	redis.InitRedis()
	seaweedfs.InitSeaWeedFS()
	dal.InitDataSource()

}

func TestFeedSpider_Capture(t *testing.T) {
	spider := &FeedSpider{}
	var sources []*model.ArticleSource
	sources = append(sources, &model.ArticleSource{Url:"http://cmsblogs.com/?feed=rss2"})
	spider.Capture(sources)
}
