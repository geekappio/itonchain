package main

import (
	"github.com/geekappio/itonchain/app/config"
	"testing"
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestWechatPublishSign(t *testing.T) {
	// Setup
	result := t.Run("Init config", func(t *testing.T) {
		config.InitAppConfig(config.DEFAULT_CONFIG_PATH)
		logging.InitLoggers()
	})
	if !result {
		log.Fatal("初始化配置失败")
		return
	}

	signature := "19f6d941cbb612d720f51f05d0b9ddb916e05a5d"
	timestamp := "1524582789"
	nonce := "4052783947"
	assert.True(t, CheckWechatPublishSign(signature, timestamp, nonce))
}
