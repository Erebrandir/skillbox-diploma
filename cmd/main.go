package main

import (
	"os"
	"os/signal"
	"skillbox-diploma/config"
	"skillbox-diploma/internal/server"
	"syscall"
)

func main() {
	config.GlobalConfig = config.GetDefaultConfig()

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
