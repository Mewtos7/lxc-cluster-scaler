package hetzner

import (
	"context"
	"log"
	"os"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

var hcloud_token string = os.Getenv("HCLOUD_TOKEN")

func init() {
	if hcloud_token == "" {
		log.Fatal("HCLOUD_TOKEN environment variable is not set")
	}
}

func CreateHetznerServer() {
	log.Println("Creating Hetzner server...")

	client := hcloud.NewClient(hcloud.WithToken(hcloud_token))
	ctx := context.TODO()

	result, _, err := client.Server.Create(ctx, hcloud.ServerCreateOpts{
		Image:    &hcloud.Image{Name: "debian-13"},
		Location: &hcloud.Location{Name: "nbg1"},
		Name:     "srvlx-nbg1-eva01",
		PublicNet: &hcloud.ServerCreatePublicNet{
			EnableIPv4: false,
			EnableIPv6: false,
			IPv4:       nil,
			IPv6:       nil,
		},
		ServerType:       &hcloud.ServerType{Name: "cpx32"},
		StartAfterCreate: hcloud.Ptr(true),
	})
	if err != nil {
		log.Printf("Error %s when creating server\n", err)
		return
	}

	err = client.Action.WaitFor(ctx, result.Action)

	server := result.Server

	log.Printf("Server created: %s\n", server.Name)
}
