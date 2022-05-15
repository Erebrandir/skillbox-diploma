package main

import (
	"os"
	"os/signal"
	"skillbox-diploma/pkg/config"
	"skillbox-diploma/pkg/server"
	"skillbox-diploma/pkg/simulator"
	"syscall"
)

func main() {
	config.GlobalConfig = config.NewConfig("config.yaml")

	go simulator.StartSimulatorServer()
	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
