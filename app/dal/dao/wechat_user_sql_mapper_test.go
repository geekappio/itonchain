package dao

import (
	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/geekappio/itonchain/app/dal"
	"testing"
	"fmt"
)

func init()  {
	config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
	logging.InitLoggers()
	dal.InitDataSource()
}

func TestSelectUser(t *testing.T) {

	wechat, _ := GetWechatUserSqlMapper(nil).SelectUser("1")

	fmt.Println(wechat);
}
