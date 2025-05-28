package main

import (
	"log"

	"github.com/arjunksofficial/lumelassignment/pkg/api"
	"github.com/arjunksofficial/lumelassignment/pkg/config"
)

func main() {
	log.Println("Starting API server...")
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}
	log.Printf("Configuration loaded: %+v", config)
	r := api.GetRouter()
	if err := r.Run(":" + config.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
