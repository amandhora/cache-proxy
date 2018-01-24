package main

import (
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_REDIS_URL           string        = "localhost:7777"
	DEFAULT_PROXY_LISTEN_PORT   int           = 8080
	DEFAULT_CACHE_CAPACITY      int           = 1000
	DEFAULT_CACHE_ENTRY_TTL     time.Duration = 600 * time.Second // 10 min
	DEFAULT_PROXY_PARALLEL_CONN int           = 1
)

type Conf struct {
	redisUrl      string
	proxyPort     int
	cacheCapacity int
	cacheExpiry   time.Duration
}

func LoadConfigParams() *Conf {

	conf := &Conf{
		redisUrl:      DEFAULT_REDIS_URL,
		proxyPort:     DEFAULT_PROXY_LISTEN_PORT,
		cacheCapacity: DEFAULT_CACHE_CAPACITY,
		cacheExpiry:   DEFAULT_CACHE_ENTRY_TTL,
	}

	redis := os.Getenv("REDIS_URL")
	if len(redis) > 0 {
		conf.redisUrl = redis
	}

	port := os.Getenv("PROXY_PORT")
	if len(port) > 0 {
		// Handle error
		conf.proxyPort, _ = strconv.Atoi(port)
	}

	cap := os.Getenv("CACHE_CAP")
	if len(cap) > 0 {
		// Handle error
		conf.cacheCapacity, _ = strconv.Atoi(cap)
	}

	ttl := os.Getenv("CACHE_TTL_SEC")
	if len(ttl) > 0 {
		// Handle error
		exp, _ := strconv.Atoi(ttl)
		conf.cacheExpiry = time.Duration(exp) * time.Second
	}

	return conf
}
