package api_delete_segment

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"segment-manager/internal/api/model"
)

type SegmentService interface {
	Delete(ctx context.Context, slug string) error
}

func New(log *slog.Logger, segmentService SegmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Failed to decode req")
			model.SendErrorResponse(w, http.StatusBadRequest, "Failed to decode req")

			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("Request validation error")
			model.SendErrorResponse(w, http.StatusBadRequest, "Request validation error")

			return
		}

		err = segmentService.Delete(r.Context(), req.Slug)
		if err != nil {
			log.Error("Failed to delete segment")
			model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete segment")

			return
		}

		render.JSON(w, r, Response{})

		return
	}
}
