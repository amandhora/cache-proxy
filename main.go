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

	InitCache(conf.cacheCapacity, conf.cacheExpiry)

	InitProxy(conf.proxyPort)

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

}
