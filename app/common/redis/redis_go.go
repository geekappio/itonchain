package redis

import (
	"log"

	"github.com/geekappio/itonchain/app/config"
	"github.com/go-redis/redis"
	"math"
)

var redisClient *redis.Client

func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.App.Redis.Address,
		DB:   config.App.Redis.DB,
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
	err := redisClient.Set(key, value, math.MaxInt64).Err()
	if err != nil {
		log.Println("redis set failed:", err)
	}
}
