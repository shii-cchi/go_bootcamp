package app

import (
	"day03/internal/config"
	"day03/internal/es_utils"
	"log"
)

func RunCreateIndexAndUploadData() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading the config: %s", err)
	}

	es_utils.CreateIndexAndUploadData(cfg)
}
