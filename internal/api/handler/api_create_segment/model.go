package api_create_segment

import (
	"segment-manager/internal/api/model"
)

type Request struct {
	Slug string `json:"slug" validate:"required"`
}

type Response struct {
	SegmentID int64              `json:"segment_id"`
	Error     *model.ResultError `json:"error,omitempty"`
}
