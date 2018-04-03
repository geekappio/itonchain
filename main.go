package main

import (
	. "common/logging"
	. "config"
	. "dao"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	. "web"
	"util"
)

func initConfig() error {

	// Parse command line arguments
	// Config file path
	configPath := flag.String("config", DEFAULT_CONFIG_PATH, "Needs config file path.")
	flag.Parse()

	var err error

	// Init application configurations.
	err = InitAppConfig(*configPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Init Loggers
	err = InitLoggers()
	if err != nil {
		log.Fatal(err)
	}
	// Init redis.
	err = InitRedis()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Init database
	err = InitDataSource()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {

	// Init application configurations.
	err := initConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	gin.SetMode(Config.RunMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()     //获得路由实例
	//添加中间件
	router.Use(Middleware)

	router.GET(ApiRequestMapping.UserRegister, UserRegister)
	util.AddPostRouter(router, ApiRequestMapping.ArticleShare, ArticleShareHandler)

	// Handle websocket
	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	router.Run(Config.Server.Address)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}
