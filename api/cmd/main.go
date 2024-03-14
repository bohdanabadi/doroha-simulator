package main

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/server"
	"github.com/bohdanabadi/Traffic-Simulation/api/broadcast"
	"github.com/bohdanabadi/Traffic-Simulation/api/cache"
	"github.com/bohdanabadi/Traffic-Simulation/api/cron"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"log"
)

//go:generate openapi-generator-cli generate -i ../api/api.yml -g go-gin-server -o ../api/generated --global-property models
func main() {
	db.ConnectDatabase()
	cache.NewCache()
	cron.SetupDBCleanupJob()
	cfg := server.LoadConfig()
	go broadcast.H.Run()
	newServer := server.NewServer(cfg)
	log.Fatal(newServer.StartServer())

}
