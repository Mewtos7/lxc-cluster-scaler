package main

import (
	"log"

	"lxcclusterscaler/provider/hetzner"
)

func CreateServer(provider string) {
	log.Println("Selecting cloud provider:", provider)
	switch provider {
	case "hetzner":
		hetzner.CreateHetznerServer()
	default:
		log.Println("Unsupported provider")
	}
}
