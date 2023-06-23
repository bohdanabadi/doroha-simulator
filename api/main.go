package main

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	server := config.NewServer(cfg)
	log.Fatal(server.StartServer())
}
