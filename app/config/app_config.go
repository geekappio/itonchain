package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang/go/src/io/ioutil"
	"github.com/ian-kent/go-log/log"

	// "gopkg.in/yaml.v2"
	"time"
	"github.com/geekappio/itonchain/app/model"
)

// App configs.
type AppConfig struct {
	// "debug"、"release"、"test"
	RunMode        string               `json:"runMode"`
	Logging        LoggingConfig        `json:"util"`
	Server         ServerConfig         `json:"server"`
	Redis          RedisConfig          `json:"redis"`
	Database       DatabaseConfig       `json:"database"`
	XormPlus       XormPlusConfig       `json:"xormPlus"`
	SeaWeedFS      SeaWeedFSConfig      `json:"seaweed"`
	GeekappPublish GeekappPublishConfig `json:"geekapp.publish"`
	// 后台管理员先使用配置文件定义
	AdminUser 	   model.AdminUser		`json:"admin.user"`
}

type GeekappPublishConfig struct {
	Token string `json:"token"`
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

// SeaWeedFS configs.
type SeaWeedFSConfig struct {
	// 上传地址 url
	SwFSMasterUrl string        `json:"SwFSMasterUrl"`
	SwFSSchema    string        `json:"SwFSSchema"`
	SwFSFilerUrls string        `json:"SwFSFilerUrl"`
	ChunkSize     int64         `json"chunkSize"`
	Duration      time.Duration `json"duration"`
	UploadAddrUrl string        `json:"uploadAddressUrl"`
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
var App *AppConfig

// Init application configs.
// FIXME 我调试了很久yaml格式的配置，但是一直没有调通，换用json格式
func InitAppConfig(configPath string) error {
	var err error

	root, _ := os.Getwd()
	fmt.Println("Working directory is :" + root)
	fmt.Println("The config file path is :" + configPath)

	// Read config from yaml file.
	configFile, err = ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
		return err
	}

	// Set config items.
	App, err = getAppConfig()
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
		return err
	}

	return nil
}

func getAppConfig() (appConfig *AppConfig, err error) {
	if App != nil {
		return App, nil
	}

	appConfig = new(AppConfig)
	err = json.Unmarshal(configFile, appConfig)
	return
}
