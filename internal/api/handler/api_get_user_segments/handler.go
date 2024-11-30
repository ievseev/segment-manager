package api_get_user_segments

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"segment-manager/internal/api/model"
	userSegmentService "segment-manager/internal/service/user_segment"
)

type UserSegmentsService interface {
	Get(ctx context.Context, userID int64) ([]userSegmentService.Segment, error)
}

type GetUserSegments struct {
	userSegmentsService UserSegmentsService
	log                 *slog.Logger
}

func New(userSegmentsService UserSegmentsService, log *slog.Logger) *GetUserSegments {
	return &GetUserSegments{userSegmentsService: userSegmentsService, log: log}
}

func (g *GetUserSegments) Handler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userSegments, err := g.userSegmentsService.Get(r.Context(), userID)
	if err != nil {
		g.log.Error("Failed to create segment")
		model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")

		return
	}

	render.JSON(w, r, convToResponse(userSegments))

	return
}

func getUserID(r *http.Request) (int64, error) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		return 0, errors.New("user_id is required")
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid user_id")
	}
	return userID, nil
}

func convToResponse(userSegments []userSegmentService.Segment) Response {
	segments := make([]Segment, 0, len(userSegments))

	for _, segment := range userSegments {
		segments = append(segments, Segment{
			SegmentID: segment.ID,
			Slug:      segment.Slug,
		})
	}

	return Response{
		Segments: segments,
	}
}
