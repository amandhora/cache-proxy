package main

import (
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_REDIS_URL          string        = "localhost:6379"
	DEFAULT_PROXY_LISTEN_PORT  int           = 8080
	DEFAULT_CACHE_CAPACITY     int           = 1000              // Max entries in cache
	DEFAULT_CACHE_ENTRY_TTL    time.Duration = 600 * time.Second // 10 min - expiry time for cache entry
	DEFAULT_PROXY_PARALLEL_REQ int           = 10                // Num of http req processed concurrently
	DEFAULT_PROXY_MAX_CONN     int           = 1000              // Max http connection being processed
)

type Config struct {
	redisUrl       string
	proxyPort      int
	cacheCapacity  int
	cacheExpiry    time.Duration
	parallelReqCnt int
	maxConn        int
}

func LoadConfigParams() *Config {

	// We can also read parallelReqCnt and maxConn from config/env.
	conf := &Config{
		redisUrl:       DEFAULT_REDIS_URL,
		proxyPort:      DEFAULT_PROXY_LISTEN_PORT,
		cacheCapacity:  DEFAULT_CACHE_CAPACITY,
		cacheExpiry:    DEFAULT_CACHE_ENTRY_TTL,
		parallelReqCnt: DEFAULT_PROXY_PARALLEL_REQ,
		maxConn:        DEFAULT_PROXY_MAX_CONN,
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
