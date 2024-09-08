package api_delete_segment

import "segment-manager/internal/api/model"

type Request struct {
	Slug string `json:"slug" validate:"required"`
}

type Response struct {
	Error *model.ResultError `json:"error,omitempty"`
}
