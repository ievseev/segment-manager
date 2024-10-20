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

type DeleteSegment struct {
	segmentService SegmentService
	log            *slog.Logger
}

func New(segmentService SegmentService, log *slog.Logger) *DeleteSegment {
	return &DeleteSegment{segmentService: segmentService, log: log}
}

func (h *DeleteSegment) Handler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.log.Error("Failed to decode req")
		model.SendErrorResponse(w, http.StatusBadRequest, "Failed to decode req")

		return
	}

	if err := validator.New().Struct(req); err != nil {
		h.log.Error("Request validation error")
		model.SendErrorResponse(w, http.StatusBadRequest, "Request validation error")

		return
	}

	err = h.segmentService.Delete(r.Context(), req.Slug)
	if err != nil {
		h.log.Error("Failed to delete segment")
		model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete segment")

		return
	}

	render.JSON(w, r, Response{})

	return
}
