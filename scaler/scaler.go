package main

import (
	"log"
)

var cloudprovider = "hetzner"

func main() {
	log.Println("Starting scaler startup...")
	InitializeDatabase()
	CreateServer(cloudprovider)
}
