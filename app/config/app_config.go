package config

import (
	"encoding/json"

	"github.com/golang/go/src/io/ioutil"
	"github.com/ian-kent/go-log/log"
)

// App configs.
type AppConfig struct {
	// "debug"、"release"、"test"
	RunMode  string         `json:"runMode"`
	Logging  LoggingConfig  `json:"logging"`
	Server   ServerConfig   `json:"server"`
	Redis    RedisConfig    `json:"redis"`
	Database DatabaseConfig `json:"database"`
	XormPlus XormPlusConfig `json:"xormPlus"`
}

type LoggingConfig struct {
	LogLevel       string `json:"logLevel"`
	DefaultLogPath string `json:"defaultLogPath"`
	ErrorLogPath   string `json:"errorLogPath"`
	ApiLogPath     string `json:"apiLogPath"`
}

// Server configs.
type ServerConfig struct {
	Address string `json:"address"`
}

// Redis configs.
type RedisConfig struct {
	Address string `json:"address"`
	DB      int    `json:"db"`
}

//  Databae configs.
type DatabaseConfig struct {
	DriverName         string `json:"driverName"`
	DatasourceName     string `json:"datasourceName"`
	MaxIdelConnections int    `json:"maxIdelConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

// xorm configs
type XormPlusConfig struct {
	XmlDirectory  string `json:"xmlDirectory"`
	StplDirectory string `json:"stplDirectory"`
}

var configFile []byte

// Export config object.
var Config *AppConfig

// Init application configs.
// FIXME 我调试了很久yaml格式的配置，但是一直没有调通，换用json格式
func InitAppConfig(configPath string) error {
	var err error

	// Read config from yaml file.
	configFile, err = ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
		return err
	}

	// Set config items.
	Config, err = getAppConfig()
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
		return err
	}

	return nil
}

func getAppConfig() (appConfig *AppConfig, err error) {
	if Config != nil {
		return Config, nil
	}

	appConfig = new(AppConfig)
	err = json.Unmarshal(configFile, appConfig)
	return
}
