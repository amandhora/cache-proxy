package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
)

var (
	// For parallel redis queries
	rPool *redis.Pool
)

func InitRedis() {

	redisUrl := os.Getenv("REDIS_URL")

	rPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisUrl)
		},
	}
}

func RedisGet(key string) (string, error) {

	conn := rPool.Get()

	err := conn.Err()
	if err != nil {
		log.Println(err)
		return "", err
	}

	rep, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}

	return redis.String(rep, err)
}
