package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	//TODO: init config - cleanenv
	cfg := config.MustLoad(".env")

	//TODO: init logger - slog
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	//TODO: init storage - postgres
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage") // TODO сделать ошибку детальнее
		os.Exit(1)                          // идти дальше смысла нет, выходим
	}

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
