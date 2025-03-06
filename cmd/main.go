package main

import (
	"log"

	"github.com/SyedAmirAli/secure-zip-vault/internal/config"
	"github.com/SyedAmirAli/secure-zip-vault/internal/routes"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup and start the server
	r := routes.SetupRouter(cfg)
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
