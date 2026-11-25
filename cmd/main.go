package main

import (
	"log"

	"lxcclusterscaler/internal/database"
	"lxcclusterscaler/internal/server"
)

var cloudprovider = "hetzner"

func main() {
	log.Println("Starting scaler startup...")
	database.InitializeDatabase()
	server.CreateServer(cloudprovider)
}
