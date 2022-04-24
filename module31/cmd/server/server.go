package main

import (
	"log"
	"os"
	"skillbox/internal/app"
)

func main() {
	//port := []string{"", ":8080"}
	//err := app.Run(port[1])
	port := os.Args
	err := app.Run(port[1])
	if err != nil {
		log.Fatal(err)
	}
}
