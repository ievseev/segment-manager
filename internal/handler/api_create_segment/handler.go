package api_create_segment

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SegmentService interface {
	CreateSegment(ctx context.Context, slug string) error
}

func New(log *slog.Logger, segmentService SegmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Failed to decode req")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error())

			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("Request validation error")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error())

			return
		}

		err = segmentService.CreateSegment(r.Context(), req.Slug)
		if err != nil {
			log.Error("Failed to create segment")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error())

			return
		}

		render.JSON(w, r, OK())

		return
	}
}
