package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log/slog"
	"net/http"
	"os"
	"segment-manager/internal/config"
	"segment-manager/internal/handler/api_create_segment"
	"segment-manager/internal/handler/api_delete_segment"
	segService "segment-manager/internal/service/segment"
	"segment-manager/internal/storage/postgres"
	"segment-manager/internal/store/segment"
)

func main() {
	// init config - cleanenv
	cfg := config.MustLoad(".env")

	// init logger - slog
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// init storage - postgres
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage:", err)
		os.Exit(1) // идти дальше смысла нет, выходим
	}

	// apply migrations
	m, err := migrate.New(
		"file://db/migrations",
		cfg.StoragePath,
	)
	if err != nil {
		log.Error("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Error("Failed to run migrations: %v", err)
	}

	log.Info("Migrations applied successfully!")

	//TODO: init router - chi (совместим с net/http)
	segmentDB := segment.New(storage)
	segmentService := segService.New(segmentDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/api/createSegment", api_create_segment.New(log, segmentService))
	router.Post("/api/deleteSegment", api_delete_segment.New(log, segmentService))

	log.Info("server starting")

	srv := &http.Server{
		Addr:    cfg.ServicePath,
		Handler: router,
		//ReadTimeout:       0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
