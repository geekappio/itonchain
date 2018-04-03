package logging

import (
	. "config"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var Logger *logrus.Logger

var ApiLogger *logrus.Logger

func initLoggerWriter(loggerPath string) (*rotatelogs.RotateLogs, error) {
	writer, err := rotatelogs.New(
		loggerPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(loggerPath),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)

	return writer, err
}

func initCommonLogger() (*logrus.Logger, error) {
	defaultWriter, err := initLoggerWriter(Config.Logging.DefaultLogPath)
	if err != nil {
		return nil, err
	}

	errorWriter, err := initLoggerWriter(Config.Logging.ErrorLogPath)
	if err != nil {
		return nil, err
	}

	hook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: defaultWriter, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  defaultWriter,
			logrus.WarnLevel:  defaultWriter,
			logrus.ErrorLevel: errorWriter,
			logrus.FatalLevel: errorWriter,
			logrus.PanicLevel: errorWriter,
		},
		nil,
	)

	logger := logrus.New()
	level, err := logrus.ParseLevel(Config.Logging.LogLevel)
	if err != nil {
		fmt.Errorf("Not specified log level, use info as default.")
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	logger.Hooks.Add(hook)
	return logger, nil
}

func initApiLogger() (*logrus.Logger, error) {
	apiWriter, err := initLoggerWriter(Config.Logging.ApiLogPath)
	if err != nil {
		return nil, err
	}

	hook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: apiWriter, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  apiWriter,
			logrus.WarnLevel:  apiWriter,
			logrus.ErrorLevel: apiWriter,
			logrus.FatalLevel: apiWriter,
			logrus.PanicLevel: apiWriter,
		},
		nil,
	)

	logger := logrus.New()
	level, err := logrus.ParseLevel(Config.Logging.LogLevel)
	if err != nil {
		fmt.Errorf("Not specified log level, use info as default.")
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)
	logger.Hooks.Add(hook)
	return logger, nil
}

// Init all loggers, this file should be called before others in main entry.
func InitLoggers() error {
	commLogger, err := initCommonLogger()
	if err != nil {
		fmt.Errorf("Failed to init common logger", err)
		return err
	}
	Logger = commLogger

	apiLogger, err := initApiLogger()
	if err != nil {
		fmt.Errorf("Failed to init api logger", err)
		return err
	}
	ApiLogger = apiLogger

	return nil
}
