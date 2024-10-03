package api_create_segment

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"segment-manager/internal/api/model"
)

type SegmentService interface {
	CreateSegment(ctx context.Context, slug string) (int64, error)
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

		segmentID, err := segmentService.CreateSegment(r.Context(), req.Slug)
		if err != nil {
			log.Error("Failed to create segment")
			model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create segment")

			return
		}

		render.JSON(w, r, Response{
			SegmentID: segmentID,
		})

		return
	}
}