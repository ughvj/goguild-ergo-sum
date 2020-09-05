package main

import (
	"log"

	"github.com/ughvj/goguild-ergo-sum/infrastructure/grpc"
)

func main() {
	if err := grpc.LaunchCommandServer(); err != nil {
		log.Fatal("Failed to establish Goguild: " + err.Error())
	}
}