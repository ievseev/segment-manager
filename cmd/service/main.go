package main

import (
	"log/slog"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"segment-manager/db/migrations"
	"segment-manager/internal/config"
)

func main() {
	// init config - cleanenv
	cfg := config.MustLoad(".env")

	// init logger - slog
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// apply migrations
	err := migrations.Run(cfg)
	if err != nil {
		log.Error("Failed to run migrations:", err)
		os.Exit(1)
	}

	// create and run server
	server, err := NewServer(cfg, log)
	if err != nil {
		log.Error("Failed to create server:", err)
		os.Exit(1)
	}

	server.Run()
}
