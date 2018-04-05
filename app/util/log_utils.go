package util

import "github.com/geekappio/itonchain/app/common/logging"

func LogError(err error) {
	if err != nil {
		logging.Logger.Error(err)
	}
}

func LogInfo(args ...interface{}) {
	logging.Logger.Info(args)
}
