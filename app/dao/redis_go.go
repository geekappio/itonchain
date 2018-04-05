package dao

import (
	"github.com/go-redis/redis"
	"log"
	. "github.com/geekappio/itonchain/app/config"
)

var redisClient *redis.Client

func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr: Config.Redis.Address,
		DB:   Config.Redis.DB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Println("redis ping failed:", err)
		return err
	} else {
		log.Println(pong, err)
	}
	return nil
}

func Get(key string) (value string) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		log.Println("redis get failed:", err)
	} else {
		log.Printf("Get mykey: %v \n", value)
	}
	return value
}

func Set(key string, value string) {
	err := redisClient.Set(key, value, 1000000).Err()
	if err != nil {
		log.Println("redis set failed:", err)
	}
}