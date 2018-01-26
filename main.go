package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	conf := LoadConfigParams()
	log.Println(conf)

	// Initialize redis connection pool
	InitRedis(conf.redisUrl)

	// Create http Job queue
	jobQueue := make(chan Job, conf.maxConn)

	// Initialize the LRU cache
	StartLruCacheHandlers(conf, jobQueue)

	// Start the proxy at last
	InitProxy(conf.proxyPort, jobQueue)

	// Catch Ctrl + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	sig := <-signalChan

	switch sig {
	case os.Interrupt:
		//handle SIGINT
	case syscall.SIGTERM:
		//handle SIGTERM
	}

	StopLruCacheHandlers()
}
