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

	InitRedis(conf.redisUrl)

	log.Println("redis done")
	jobQueue := make(chan Job, conf.parallelReqCnt)
	StartCacheHandlers(conf, jobQueue)

	log.Println("cache done")
	InitProxy(conf.proxyPort, jobQueue)

	log.Println("initproxy done")
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

	StopCacheHandlers()
	log.Println("done")
}
