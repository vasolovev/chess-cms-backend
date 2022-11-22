package main

import (
	"log"

	"github.com/vasolovev/ChessCMS/config"
	"github.com/vasolovev/ChessCMS/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
