package main

import (
	"log"

	"github.com/ddr4869/k8s-kafka/config"
	"github.com/ddr4869/k8s-kafka/internal"
)

func main() {
	cfg := config.Init()
	server, err := internal.NewServerSetUp(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	server.Start()
}
