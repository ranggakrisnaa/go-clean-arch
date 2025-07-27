package main

import (
	"log"

	"github.com/ranggakrisna/go-clean-arch/config"
	"github.com/ranggakrisna/go-clean-arch/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Configuration
	_ = godotenv.Load(".env")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
