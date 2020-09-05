package main

import (
	"log"
	"fmt"

	"github.com/ughvj/goguild-ergo-sum/infrastructure/grpc"
)

func main() {
	fmt.Println("** Goguild ergo sum is established")
	if err := grpc.LaunchCommandServer(); err != nil {
		log.Fatal("Failed to establish Goguild: " + err.Error())
	}
}