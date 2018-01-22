package main

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

var (
	// For parallel redis queries
	rPool *redis.Pool
)

func InitRedis(addr string, port int) {

	rPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr+":"+strconv.Itoa(port))
		},
	}
}

func RedisGet(key string) (string, error) {

	conn := rPool.Get()

	err := conn.Err()
	if err != nil {
		return "", err
	}

	rep, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}

	return redis.String(rep, err)
}
