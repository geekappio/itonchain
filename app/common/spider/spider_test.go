package spider

import (
	"testing"

	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/dal"
	"github.com/geekappio/itonchain/app/dal/entity"
	"time"
)

func init()  {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	redis.InitRedis()
	seaweedfs.InitSeaWeedFS()
	dal.InitDataSource()
}

func TestFeedSpider_Capture(t *testing.T) {
	sources := make([]*entity.ArticleSource, 1)
	sources[0] = &entity.ArticleSource{
		SourceType: "FEED",
		SourceUrl: "https://wanqu.co/feed/",
	}
	Capture(sources)
	time.Sleep(10 * time.Minute)
}
