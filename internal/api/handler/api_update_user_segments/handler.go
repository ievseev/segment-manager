package api_update_user_segments

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"segment-manager/internal/api/model"
)

type UserSegmentsService interface {
	Update(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error
}

type CreateSegment struct {
	segmentService UserSegmentsService
	log            *slog.Logger
}

func New(segmentService UserSegmentsService, log *slog.Logger) *CreateSegment {
	return &CreateSegment{
		segmentService: segmentService,
		log:            log,
	}
}

func (h *CreateSegment) Handler(w http.ResponseWriter, r *http.Request) {
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

	err = h.segmentService.Update(r.Context(), req.UserID, req.SlugsToAdd, req.SlugsToDelete)
	if err != nil {
		h.log.Error("Failed to create segment")
		model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create segment")

		return
	}

	render.JSON(w, r, Response{})

	return
}
