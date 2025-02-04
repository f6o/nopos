package main

import (
	"log"

	"github.com/f6o/napos/server"
)

func main() {
	server := server.SimpleDealerServer{}
	err := server.StartServer()
	if err != nil {
		log.Fatalln(err)
	}
}
