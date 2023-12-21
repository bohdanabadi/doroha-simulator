package main

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/server"
	"github.com/bohdanabadi/Traffic-Simulation/api/broadcast"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"github.com/gorilla/websocket"
	"log"
)

var conn *websocket.Conn

func main() {
	db.ConnectDatabase()
	go broadcast.H.Run()
	cfg := server.LoadConfig()
	newServer := server.NewServer(cfg)
	log.Fatal(newServer.StartServer())
}
