package main

import (
	"time"
)

const (
	DEFAULT_CACHE_CAPACITY      int           = 10
	DEFAULT_CACHE_ENTRY_TTL     time.Duration = 10 * time.Second // 10 sec
	DEFAULT_PROXY_LISTEN_PORT   int           = 8080
	DEFAULT_PROXY_PARALLEL_CONN int           = 1
	DEFAULT_REDIS_SERVER_HOST   string        = "localhost"
	DEFAULT_REDIS_SERVER_PORT   int           = 7777
)
