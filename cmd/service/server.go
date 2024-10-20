package main

import (
	"log/slog"
	"net/http"

	"segment-manager/internal/api/handler/api_create_segment"
	"segment-manager/internal/api/handler/api_create_user"
	"segment-manager/internal/api/handler/api_delete_segment"
	"segment-manager/internal/config"
	segService "segment-manager/internal/service/segment"
	usService "segment-manager/internal/service/user"
	"segment-manager/internal/storage/postgres"
	"segment-manager/internal/store/segment"
	"segment-manager/internal/store/user"

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

	// init services
	segmentDB := segment.New(storage)
	userDB := user.New(storage)
	segmentService := segService.New(segmentDB)
	userService := usService.New(userDB)

	// init router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/api/createSegment", api_create_segment.New(segmentService, log).Handler)
	router.Post("/api/deleteSegment", api_delete_segment.New(segmentService, log).Handler)
	router.Post("/api/createUser", api_create_user.New(userService, log).Handler)

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
