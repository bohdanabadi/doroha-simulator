package main

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/server"
	"github.com/bohdanabadi/Traffic-Simulation/api/broadcast"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"log"
)

//go:generate openapi-generator-cli generate -i ../api/api.yml -g go-gin-server -o ../api/generated --global-property models
func main() {
	db.ConnectDatabase()
	go broadcast.H.Run()
	cfg := server.LoadConfig()
	newServer := server.NewServer(cfg)
	log.Fatal(newServer.StartServer())

}
