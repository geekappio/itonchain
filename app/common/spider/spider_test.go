package spider

import (
	"testing"

	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/common/redis"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
	"github.com/geekappio/itonchain/app/dal"
	"time"
	"github.com/geekappio/itonchain/app/dal/dao"
)

func init()  {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	redis.InitRedis()
	seaweedfs.InitSeaWeedFS()
	dal.InitDataSource()
}

const PAGE_SIZE = 10

func TestFeedSpider_Capture(t *testing.T) {
	articleSourceSqlMapper := dao.GetArticleSourceSqlMapper(nil)

	total, _ := articleSourceSqlMapper.CountArticleSources()
	for i := 1; i <= (total+PAGE_SIZE-1) / PAGE_SIZE; i++ {
		sources, _ := articleSourceSqlMapper.SelectArticleSources(i, PAGE_SIZE)
		Capture(sources)
	}
	time.Sleep(10 * time.Minute)
}
