package main

import (
	"os"
	"os/signal"
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/server"
	"syscall"
)

func main() {
	config.GlobalConfig = config.NewConfig("config/config.yaml")

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
