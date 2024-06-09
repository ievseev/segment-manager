package main

import (
	"log/slog"
	"net/http"

	"segment-manager/internal/config"
	"segment-manager/internal/handler/api_create_segment"
	"segment-manager/internal/handler/api_delete_segment"
	segService "segment-manager/internal/service/segment"
	"segment-manager/internal/storage/postgres"
	"segment-manager/internal/store/segment"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	cfg    *config.Config
	log    *slog.Logger
	router *chi.Mux
}

func NewServer(cfg *config.Config, log *slog.Logger) (*Server, error) {
	// init storage - postgres
	storage, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	segmentDB := segment.New(storage)
	segmentService := segService.New(segmentDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/api/createSegment", api_create_segment.New(log, segmentService))
	router.Post("/api/deleteSegment", api_delete_segment.New(log, segmentService))

	return &Server{
		cfg:    cfg,
		log:    log,
		router: router,
	}, nil
}

func (s *Server) Run() {
	s.log.Info("server starting")

	srv := &http.Server{
		Addr:    s.cfg.ServicePath,
		Handler: s.router,
		//ReadTimeout:       0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
	}

	if err := srv.ListenAndServe(); err != nil {
		s.log.Error("failed to start server")
	}

	s.log.Error("server stopped")
}
