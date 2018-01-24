package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {

	InitRedis()

	NewCache(DEFAULT_CACHE_CAPACITY)

	InitProxy(DEFAULT_PROXY_LISTEN_PORT)

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
