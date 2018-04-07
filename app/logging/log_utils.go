package logging

import (
	"github.com/geekappio/itonchain/app/common/logging"
	"github.com/json-iterator/go"
)

func LogDebug(args ...interface{}) {
	logging.Logger.Info(args)
}

func LogInfo(args ...interface{}) {
	logging.Logger.Info(args)
}

func LogWarn(args ...interface{}) {
	logging.Logger.Info(args)
}

func LogError(args ...interface{}) {
	logging.Logger.Error(args)
}

func LogRequest(requestUri string, requestData []byte) {
	logging.ApiLogger.Info("API: ", requestUri, ", Request Data: ", string(requestData))
}

func LogResponse(requestUri string, responseModel *interface{}) {
	responseStr, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(responseModel)
	if err != nil {
		LogError("Error happened when decoding service all result.", err)
	}
	logging.ApiLogger.Info("API: ", requestUri, ", Response Data: ", responseStr)
}
