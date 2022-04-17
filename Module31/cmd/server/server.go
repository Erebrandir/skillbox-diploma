package main

import (
	"log"
	"skillbox/internal/app"
)

func main() {
	port := []string{"", ":8080"}
	err := app.Run(port[0])
	if err != nil {
		log.Fatal(err)
	}

}
